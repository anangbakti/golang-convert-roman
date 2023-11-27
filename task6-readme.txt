Execution Flow of the Program:

- The first step is mapping a collection of key-value pairs of Roman numerals and numbers.
- The program will then read the roman.txt file and loop through each line of Roman numerals.
- For each Roman numeral, it becomes the input for the convertToNumber function, 
  where this function will return a result in the form of a number.
- The convertToNumber function will take the value of the last numeral.
- Then, looping backward through the Roman numeral, to add or subtract the result, adjusting its front numeral.
- If the last numeral is less than or equal to the front numeral, the result value = result + front numeral.
- If the last numeral is greater than the front numeral, the result value = result - front numeral.

How to Run the Program:

- Make sure Golang is installed.
- Open the terminal in the folder containing main.go.
- Place roman.txt in the same folder as main.go.
- In the terminal, type: go run main.go.
- The program will output Roman numerals along with their corresponding numbers.

Reference : https://medium.com/@anuragsahani0123/roman-to-integer-solution-in-golang-22e156ebe7f9

package main
import "fmt"
func main() {
fmt.Println(romanToInt("LVI"))
}
//Need to Know How to find the last element of String or String x.
//1. For finding Last element you need to know how we calculate the total length of string.
//solution:
//we are usgin len(string_name) for finding length.
// for finding last element 
// string_name[len(string_name)-1:len(string_name)]
// eg: str :="abcd"
// str[4-1 : 4]
// we have passed [3:4] to the slice expression. so it starts the extraction at position 3 and ends at position 4 (which is excluded).
//2. For finding second last element 
//   string_name[ len(string_name) - 2 : len(string_name) - 1 ]
// eg: str :="abcd"
// str[4-2 : 3]
// we have passed [2:3] to the slice expression. so it starts the extraction at position 2 and ends at position 3 (which is excluded).
// you Need to know the use of map in go lang

func romanToInt(s string) int {
know := map[string]int{
            "I": 1,
            "V": 5,
            "X": 10,
            "L": 50,
            "C": 100,
            "D": 500,
            "M": 1000,
}
lengthOfString := len(s)
lastElement := s[len(s)-1 : lengthOfString]
var result int
//Here we getting the int value from map
result = know[lastElement]
// Iterating the loop from right to left
for i := len(s) - 1; i > 0; i-- {
/fmt.Println("Knowing Value : last ::-", know[s[i:i+1]])
//fmt.Println("Knowing Value : 2nd Last ::-", know[s[i-1:i]])
//Here we are checking if last value is less then or equal to second last value of string we add the value and viceversa
    if know[s[i:i+1]] <= know[s[i-1:i]] {
       result += know[s[i-1:i]]
    } else {
       result -= know[s[i-1:i]]
    }
  }
return result
}

Roman to number rules from : https://www.thevbprogrammer.com/Ch08/08-10-RomanNumerals.htm

'--------------------------------------------------------------------------------------
Private Function ValidRomanInput(ByVal pstrRN As String, ByRef pstrMsg As String) _
As Boolean
'--------------------------------------------------------------------------------------
        
    ValidRomanInput = False         ' Guilty until proven innocent!
        
    ' 'D', 'L', or 'V' may only appear at most once in the string
    If GetSubstringCount(pstrRN, "D") > 1 _
    Or GetSubstringCount(pstrRN, "L") > 1 _
    Or GetSubstringCount(pstrRN, "V") > 1 Then
        pstrMsg = "'D', 'L', or 'V' may only appear at most once."
        Exit Function
    End If
    
    ' no more than 3 consecutive Ms, Cs, Xs or Is:
    If InStr(pstrRN, "MMMM") > 0 _
    Or InStr(pstrRN, "CCCC") > 0 _
    Or InStr(pstrRN, "XXXX") > 0 _
    Or InStr(pstrRN, "IIII") > 0 Then
        pstrMsg = "'M', 'C', 'X', or 'I' may appear no more than three times in a row."
        Exit Function
    End If
    
    ' Outright illegal sequences:
    ' --------------------------
    ' Only I, X, and C can be used for subtraction (V, L, and D cannot). Therefore, the
    ' following pairs of letter are invalid: VX, VL, VC, VD, VM, LC, LD, LM, DM.
    ' When subtracting, the value of the letter being subtracted from cannot be more than
    ' 10 times the value of letter being used for subtraction. Therefore, the following pairs
    ' of letters are invalid: IL, IC, ID, IM, XD, XM.
    If InStr(pstrRN, "IL") > 0 _
    Or InStr(pstrRN, "IC") > 0 _
    Or InStr(pstrRN, "ID") > 0 _
    Or InStr(pstrRN, "IM") > 0 _
    Or InStr(pstrRN, "XD") > 0 _
    Or InStr(pstrRN, "XM") > 0 _
    Or InStr(pstrRN, "VX") > 0 _
    Or InStr(pstrRN, "VL") > 0 _
    Or InStr(pstrRN, "VC") > 0 _
    Or InStr(pstrRN, "VD") > 0 _
    Or InStr(pstrRN, "VM") > 0 _
    Or InStr(pstrRN, "LC") > 0 _
    Or InStr(pstrRN, "LD") > 0 _
    Or InStr(pstrRN, "LM") > 0 _
    Or InStr(pstrRN, "DM") > 0 _
    Then
        pstrMsg = "The Roman Numeral string contains an illegal sequence of characters."
        Exit Function
    End If
    
    ' Other illegal sequences:

    ' Once a letter has been used as a subtraction modifier, that letter cannot appear again
    ' in the string, unless that letter itself is subtracted from. For example, CDC is not
    ' valid (you would be subtracting 100 from 500, then adding it right back) -
    ' but CDXC (for 490) is valid. Similarly, XCX is not valid, but XCIX is.
    ' To summarize:
    ' C cannot follow CM or CD except in case of XC.
    ' X cannot follow XC or XL except in the case of IX.
    If AFollowsBInCExceptD("C", "CD", pstrRN, "XC") Then pstrMsg = "'C' cannot follow 'CD' except for 'XC'.": Exit Function
    If AFollowsBInCExceptD("C", "CM", pstrRN, "XC") Then pstrMsg = "'C' cannot follow 'CD' except for 'XC'.": Exit Function
    If AFollowsBInCExceptD("X", "XL", pstrRN, "IX") Then pstrMsg = "'X' cannot follow 'XL' except for 'IX'.": Exit Function
    If AFollowsBInCExceptD("X", "XC", pstrRN, "IX") Then pstrMsg = "'X' cannot follow 'XL' except for 'IX'.": Exit Function
    If AFollowsBInC("I", "IV", pstrRN) Then pstrMsg = "'I' cannot follow 'IV'.": Exit Function
    If AFollowsBInC("I", "IX", pstrRN) Then pstrMsg = "'I' cannot follow 'IV'.": Exit Function
    
    ' Once a letter has been subtracted from, neither it nor the next lowest multiple of 5 may
    ' appear again in the string - so neither X nor V can follow IX, neither C nor L may follow
    ' XC, and neither M nor D may follow CM.
    If AFollowsBInC("X", "IX", pstrRN) Then pstrMsg = "'X' cannot follow 'IX'.": Exit Function
    If AFollowsBInC("V", "IX", pstrRN) Then pstrMsg = "'V' cannot follow 'IX'.": Exit Function
    If AFollowsBInC("C", "XC", pstrRN) Then pstrMsg = "'C' cannot follow 'XC'.": Exit Function
    If AFollowsBInC("L", "XC", pstrRN) Then pstrMsg = "'L' cannot follow 'XC'.": Exit Function
    If AFollowsBInC("M", "CM", pstrRN) Then pstrMsg = "'M' cannot follow 'CM'.": Exit Function
    If AFollowsBInC("D", "CM", pstrRN) Then pstrMsg = "'D' cannot follow 'CM'.": Exit Function
    
    '  A letter cannot be used as a subtraction modifier if that letter, or the next highest
    ' multiple of 5, appears previously in the string - so IV or IX cannot follow I or V,
    ' XL or XC cannot follow X or L, and CD or CM cannot follow C or D.
    If AFollowsBInC("IV", "I", pstrRN) Then pstrMsg = "'IV' cannot follow 'I'.": Exit Function
    If AFollowsBInC("IX", "I", pstrRN) Then pstrMsg = "'IX' cannot follow 'I'.": Exit Function
    If AFollowsBInC("IX", "V", pstrRN) Then pstrMsg = "'IX' cannot follow 'V'.": Exit Function
    If AFollowsBInC("XL", "X", pstrRN) Then pstrMsg = "'XL' cannot follow 'X'.": Exit Function
    If AFollowsBInC("XC", "X", pstrRN) Then pstrMsg = "'XC' cannot follow 'X'.": Exit Function
    If AFollowsBInC("XC", "L", pstrRN) Then pstrMsg = "'XC' cannot follow 'L'.": Exit Function
    If AFollowsBInC("CD", "C", pstrRN) Then pstrMsg = "'CD' cannot follow 'C'.": Exit Function
    If AFollowsBInC("CM", "C", pstrRN) Then pstrMsg = "'CM' cannot follow 'C'.": Exit Function
    If AFollowsBInC("CM", "D", pstrRN) Then pstrMsg = "'CM' cannot follow 'D'.": Exit Function
    
    ' If we make it here, the Roman Numeral input was valid ...
    ValidRomanInput = True

End Function

'--------------------------------------------------------------------------------------
Private Function GetSubstringCount(ByVal pstrMainString As String, _
                                   ByVal pstrSubstring As String) _
As Long
'--------------------------------------------------------------------------------------
    
    Dim lngX As Long
    Dim lngY As Long
    
    If pstrMainString = "" Then
        GetSubstringCount = 0
    Else
        lngX = InStr(1, pstrMainString, pstrSubstring, vbBinaryCompare)
        If lngX = 0 Then
            GetSubstringCount = 0
        Else
            lngX = 0
            For lngY = 1 To Len(pstrMainString)
                If Mid$(pstrMainString, lngY, Len(pstrSubstring)) = pstrSubstring Then
                    lngX = lngX + 1
                End If
            Next lngY
            GetSubstringCount = lngX
        End If
    End If

End Function

'--------------------------------------------------------------------------------------
Private Function AFollowsBInC(pstrA As String, pstrB As String, pstrC As String) As Boolean
'--------------------------------------------------------------------------------------
    
    Dim lngTestPos  As Long
    
    lngTestPos = InStr(pstrC, pstrB)
    
    If lngTestPos > 0 Then
        If InStr(lngTestPos + Len(pstrB), pstrC, pstrA, vbTextCompare) Then
            AFollowsBInC = True
        Else
            AFollowsBInC = False
        End If
    Else
        AFollowsBInC = False
    End If
        

End Function

'--------------------------------------------------------------------------------------
Private Function AFollowsBInCExceptD(pstrA As String, _
                                     pstrB As String, _
                                     pstrC As String, _
                                     pstrD As String) _
As Boolean
'--------------------------------------------------------------------------------------
    
    Dim lngTestPos  As Long
    Dim strRemChars As String
    
    lngTestPos = InStr(pstrC, pstrB)
    
    If lngTestPos = 0 Then
        AFollowsBInCExceptD = False
        Exit Function
    End If
    
    strRemChars = Mid$(pstrC, lngTestPos + Len(pstrB))
    strRemChars = Replace$(strRemChars, pstrD, "")
    If InStr(strRemChars, pstrA) > 0 Then
        AFollowsBInCExceptD = True
    Else
        AFollowsBInCExceptD = False
    End If

End Function
