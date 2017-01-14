---
title: SWGEmu - .Net Stuffz
published: true
date: 2008-05-29 23:08:00 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/dotnet/archive/2008/05/29/swgemu-net-stuffz.aspx
file: swgemu-net-stuffz.aspx
path: /blogs/dotnet/archive/2008/05/29/
author: tom anderson
words: 7840
---
So a while back some of you may know that I was writing a pure .Net emulator.  That went "ok" and I had made zone in both pre-cu and nge based clients.  But after a while, the time constraints just couldn't be dealt with, and the project was, how do I say, huge?

![][1]

One of the big issues that I had when writing the emulator in .Net was the encryption, decryption, and the crc calculations.  Originally I had just used a c api to pull it off.  This bothered me literally for some time as I just wanted a pure .Net solution.

So today, I spent a bit of time looking at the great documentation on the encryption/decryption algorigthm that the SWGEmu team has put up (if you recall, I used to be a part of that team for a very short lived time).

Below is the test app that I created and played around with to get these things working in pure (you got it) VB.Net.  I know the C# port would have been WAY easier using unsafe or even almost just duplicating the code, but I wanted a direct port that was spot on.

 
    
    
       1:  Public Class Program
    
    
       2:   
    
    
       3:      Public Shared Sub Main(ByVal args() As String)
    
    
       4:          Dim rawArray() As Byte = New Byte() {0, 9, 0, 1, 0, 2, &HAB, &H43, &HE3, &HD5, 0, &HFF, 0, &H11, &H45, &H32, &H76, &H43, &HD4, &HF1, 0, &HAB, &HCD}
    
    
       5:          Dim encryptedArray() As Byte
    
    
       6:          Dim decryptedArray() As Byte
    
    
       7:          Dim crcSeed As Integer = &H1D4E3287
    
    
       8:   
    
    
       9:          Console.WriteLine("Raw Bytes:")
    
    
      10:          For Each bt As Byte In rawArray
    
    
      11:              Console.Write("{0:X2} ", bt)
    
    
      12:          Next
    
    
      13:          Console.WriteLine()
    
    
      14:          Console.WriteLine()
    
    
      15:   
    
    
      16:   
    
    
      17:          encryptedArray = Encrypt(Strip(rawArray), crcSeed)
    
    
      18:          Console.WriteLine("Encrypted Bytes:")
    
    
      19:          For Each bt As Byte In encryptedArray
    
    
      20:              Console.Write("{0:X2} ", bt)
    
    
      21:          Next
    
    
      22:          Console.WriteLine()
    
    
      23:          Console.WriteLine()
    
    
      24:   
    
    
      25:          decryptedArray = Decrypt(encryptedArray, crcSeed)
    
    
      26:          Console.WriteLine("Decrypted Bytes:")
    
    
      27:          For Each bt As Byte In decryptedArray
    
    
      28:              Console.Write("{0:X2} ", bt)
    
    
      29:          Next
    
    
      30:          Console.WriteLine()
    
    
      31:          Console.WriteLine()
    
    
      32:   
    
    
      33:          Console.WriteLine("CRC Test")
    
    
      34:          Dim crc As Integer = CRC32(New Byte() {&H0, &H5, &HAA, &HBB, &HCC, &HDD, &H0, &H6, &H0}, 9, crcSeed)
    
    
      35:          Console.WriteLine("My CRC32:  {0:X4}", crc)
    
    
      36:          crc = API.GenerateCrc(New Byte() {&H0, &H5, &HAA, &HBB, &HCC, &HDD, &H0, &H6, &H0}, crcSeed)
    
    
      37:          Console.WriteLine("C   CRC32: {0:X4}", crc)
    
    
      38:   
    
    
      39:          Console.ReadLine()
    
    
      40:      End Sub
    
    
      41:   
    
    
      42:      Private Shared Function Strip(ByVal inData() As Byte) As Byte()
    
    
      43:          Dim retVal As New List(Of Byte)
    
    
      44:          retVal.AddRange(inData)
    
    
      45:   
    
    
      46:          'strip channel and op code
    
    
      47:          If retVal(0) = 0 Then
    
    
      48:              retVal.RemoveAt(0)
    
    
      49:              retVal.RemoveAt(0)
    
    
      50:          End If
    
    
      51:   
    
    
      52:          'strip crc
    
    
      53:          retVal.RemoveAt(retVal.Count - 1)
    
    
      54:          retVal.RemoveAt(retVal.Count - 1)
    
    
      55:   
    
    
      56:          Return retVal.ToArray
    
    
      57:      End Function
    
    
      58:   
    
    
      59:      Public Shared Function Encrypt(ByVal inData() As Byte, ByVal seed As Integer) As Byte()
    
    
      60:          Dim count As Integer = Math.Floor(Decimal.Parse(inData.Length) / 4D)
    
    
      61:          Dim remainder As Integer = inData.Length - (count * 4)
    
    
      62:          Dim seedBytes() As Byte = BitConverter.GetBytes(seed)
    
    
      63:   
    
    
      64:          Array.Reverse(seedBytes)
    
    
      65:   
    
    
      66:          For i As Integer = 0 To count - 1
    
    
      67:              For b As Integer = 0 To 3
    
    
      68:                  inData((i * 4)   b) = inData((i * 4)   b) Xor seedBytes(b)
    
    
      69:              Next
    
    
      70:              For b As Integer = 0 To 3
    
    
      71:                  seedBytes(b) = inData((i * 4)   b)
    
    
      72:              Next
    
    
      73:          Next
    
    
      74:   
    
    
      75:          For i As Integer = 0 To remainder - 1
    
    
      76:              inData((count * 4)   i) = inData((count * 4)   i) Xor seedBytes(0)
    
    
      77:          Next
    
    
      78:   
    
    
      79:          Return inData
    
    
      80:      End Function
    
    
      81:   
    
    
      82:      Public Shared Function Decrypt(ByVal inData() As Byte, ByVal seed As Integer) As Byte()
    
    
      83:          Dim count As Integer = Math.Floor(Decimal.Parse(inData.Length) / 4D)
    
    
      84:          Dim remainder As Integer = inData.Length - (count * 4)
    
    
      85:          Dim seedBytes() As Byte = BitConverter.GetBytes(seed)
    
    
      86:   
    
    
      87:          Array.Reverse(seedBytes)
    
    
      88:   
    
    
      89:          For i As Integer = 0 To count - 1
    
    
      90:              Dim newSeed(3) As Byte
    
    
      91:   
    
    
      92:              For b As Integer = 0 To 3
    
    
      93:                  newSeed(b) = inData((i * 4)   b)
    
    
      94:              Next
    
    
      95:              For b As Integer = 0 To 3
    
    
      96:                  inData((i * 4)   b) = inData((i * 4)   b) Xor seedBytes(b)
    
    
      97:              Next
    
    
      98:              For b As Integer = 0 To 3
    
    
      99:                  seedBytes(b) = newSeed(b)
    
    
     100:              Next
    
    
     101:          Next
    
    
     102:   
    
    
     103:          For i As Integer = 0 To remainder - 1
    
    
     104:              inData((count * 4)   i) = inData((count * 4)   i) Xor seedBytes(0)
    
    
     105:          Next
    
    
     106:   
    
    
     107:          Return inData
    
    
     108:      End Function
    
    
     109:   
    
    
     110:      Public Shared Function CRC32(ByVal inData() As Byte, ByVal length As Integer, ByVal nCrcSeed As Integer) As Integer
    
    
     111:          'unsigned int nCrc = g_nCrcTable[(~nCrcSeed) & 0xFF];
    
    
     112:          Dim nCrc As UInt32 = m_CRCTable(Not nCrcSeed And &HFF)
    
    
     113:   
    
    
     114:          'nCrc ^= 0x00FFFFFF;
    
    
     115:          nCrc = nCrc Xor &HFFFFFF
    
    
     116:   
    
    
     117:          'unsigned int nIndex = (nCrcSeed >> 8) ^ nCrc;
    
    
     118:          Dim nIndex As UInt32 = (nCrcSeed >> 8) Xor nCrc
    
    
     119:   
    
    
     120:          'nCrc = (nCrc >> 8) & 0x00FFFFFF;
    
    
     121:          nCrc = (nCrc >> 8) And &HFFFFFF
    
    
     122:   
    
    
     123:          'nCrc ^= g_nCrcTable[nIndex & 0xFF];
    
    
     124:          nCrc = nCrc Xor m_CRCTable(nIndex And &HFF)
    
    
     125:   
    
    
     126:          'nIndex = (nCrcSeed >> 16) ^ nCrc;
    
    
     127:          nIndex = (nCrcSeed >> 16) Xor nCrc
    
    
     128:   
    
    
     129:          'nCrc = (nCrc >> 8) & 0x00FFFFFF;
    
    
     130:          nCrc = (nCrc >> 8) And &HFFFFFF
    
    
     131:   
    
    
     132:          'nCrc ^= g_nCrcTable[nIndex & 0xFF];
    
    
     133:          nCrc = nCrc Xor m_CRCTable(nIndex And &HFF)
    
    
     134:   
    
    
     135:          'nIndex = (nCrcSeed >> 24) ^ nCrc;
    
    
     136:          nIndex = (nCrcSeed >> 24) Xor nCrc
    
    
     137:   
    
    
     138:          'nCrc = (nCrc >> 8) &0x00FFFFFF;
    
    
     139:          nCrc = (nCrc >> 8) And &HFFFFFF
    
    
     140:   
    
    
     141:          'nCrc ^= g_nCrcTable[nIndex & 0xFF];
    
    
     142:          nCrc = nCrc Xor m_CRCTable(nIndex And &HFF)
    
    
     143:   
    
    
     144:          'for( short i = 0; i < nLength; i   ) {
    
    
     145:          For i As Integer = 0 To length - 1
    
    
     146:              'nIndex = (pData[ i]) ^ nCrc;
    
    
     147:              nIndex = inData(i) Xor nCrc
    
    
     148:              'nCrc = (nCrc >> 8) & 0x00FFFFFF;
    
    
     149:              nCrc = (nCrc >> 8) And &HFFFFFF
    
    
     150:              'nCrc ^= g_nCrcTable[nIndex & 0xFF];
    
    
     151:              nCrc = nCrc Xor m_CRCTable(nIndex And &HFF)
    
    
     152:   
    
    
     153:          Next
    
    
     154:          '}
    
    
     155:   
    
    
     156:          'return ~nCrc;
    
    
     157:          Return Not nCrc
    
    
     158:      End Function
    
    
     159:   
    
    
     160:      '   void CRC_(unsigned char *data, int size, unsigned long &crc){
    
    
     161:      '        if (!size) return;
    
    
     162:      '        for (int i = 0; i < size; i  )
    
    
     163:      '            crc = (crc >> 8) ^ crc_table[(crc & 0xff) ^ data[ i]];
    
    
     164:      '   }
    
    
     165:   
    
    
     166:      Private Shared m_CRCTable() As UInt32 = New UInt32() { _
    
    
     167:                                                      &H0UL, &H77073096UL, &HEE0E612CUL, &H990951BAUL, &H76DC419UL, &H706AF48FUL, _
    
    
     168:                                                      &HE963A535UL, &H9E6495A3UL, &HEDB8832UL, &H79DCB8A4UL, &HE0D5E91EUL, &H97D2D988UL, _
    
    
     169:                                                      &H9B64C2BUL, &H7EB17CBDUL, &HE7B82D07UL, &H90BF1D91UL, &H1DB71064UL, &H6AB020F2UL, _
    
    
     170:                                                      &HF3B97148UL, &H84BE41DEUL, &H1ADAD47DUL, &H6DDDE4EBUL, &HF4D4B551UL, &H83D385C7UL, _
    
    
     171:                                                      &H136C9856UL, &H646BA8C0UL, &HFD62F97AUL, &H8A65C9ECUL, &H14015C4FUL, &H63066CD9UL, _
    
    
     172:                                                      &HFA0F3D63UL, &H8D080DF5UL, &H3B6E20C8UL, &H4C69105EUL, &HD56041E4UL, &HA2677172UL, _
    
    
     173:                                                      &H3C03E4D1UL, &H4B04D447UL, &HD20D85FDUL, &HA50AB56BUL, &H35B5A8FAUL, &H42B2986CUL, _
    
    
     174:                                                      &HDBBBC9D6UL, &HACBCF940UL, &H32D86CE3UL, &H45DF5C75UL, &HDCD60DCFUL, &HABD13D59UL, _
    
    
     175:                                                      &H26D930ACUL, &H51DE003AUL, &HC8D75180UL, &HBFD06116UL, &H21B4F4B5UL, &H56B3C423UL, _
    
    
     176:                                                      &HCFBA9599UL, &HB8BDA50FUL, &H2802B89EUL, &H5F058808UL, &HC60CD9B2UL, &HB10BE924UL, _
    
    
     177:                                                      &H2F6F7C87UL, &H58684C11UL, &HC1611DABUL, &HB6662D3DUL, &H76DC4190UL, &H1DB7106UL, _
    
    
     178:                                                      &H98D220BCUL, &HEFD5102AUL, &H71B18589UL, &H6B6B51FUL, &H9FBFE4A5UL, &HE8B8D433UL, _
    
    
     179:                                                      &H7807C9A2UL, &HF00F934UL, &H9609A88EUL, &HE10E9818UL, &H7F6A0DBBUL, &H86D3D2DUL, _
    
    
     180:                                                      &H91646C97UL, &HE6635C01UL, &H6B6B51F4UL, &H1C6C6162UL, &H856530D8UL, &HF262004EUL, _
    
    
     181:                                                      &H6C0695EDUL, &H1B01A57BUL, &H8208F4C1UL, &HF50FC457UL, &H65B0D9C6UL, &H12B7E950UL, _
    
    
     182:                                                      &H8BBEB8EAUL, &HFCB9887CUL, &H62DD1DDFUL, &H15DA2D49UL, &H8CD37CF3UL, &HFBD44C65UL, _
    
    
     183:                                                      &H4DB26158UL, &H3AB551CEUL, &HA3BC0074UL, &HD4BB30E2UL, &H4ADFA541UL, &H3DD895D7UL, _
    
    
     184:                                                      &HA4D1C46DUL, &HD3D6F4FBUL, &H4369E96AUL, &H346ED9FCUL, &HAD678846UL, &HDA60B8D0UL, _
    
    
     185:                                                      &H44042D73UL, &H33031DE5UL, &HAA0A4C5FUL, &HDD0D7CC9UL, &H5005713CUL, &H270241AAUL, _
    
    
     186:                                                      &HBE0B1010UL, &HC90C2086UL, &H5768B525UL, &H206F85B3UL, &HB966D409UL, &HCE61E49FUL, _
    
    
     187:                                                      &H5EDEF90EUL, &H29D9C998UL, &HB0D09822UL, &HC7D7A8B4UL, &H59B33D17UL, &H2EB40D81UL, _
    
    
     188:                                                      &HB7BD5C3BUL, &HC0BA6CADUL, &HEDB88320UL, &H9ABFB3B6UL, &H3B6E20CUL, &H74B1D29AUL, _
    
    
     189:                                                      &HEAD54739UL, &H9DD277AFUL, &H4DB2615UL, &H73DC1683UL, &HE3630B12UL, &H94643B84UL, _
    
    
     190:                                                      &HD6D6A3EUL, &H7A6A5AA8UL, &HE40ECF0BUL, &H9309FF9DUL, &HA00AE27UL, &H7D079EB1UL, _
    
    
     191:                                                      &HF00F9344UL, &H8708A3D2UL, &H1E01F268UL, &H6906C2FEUL, &HF762575DUL, &H806567CBUL, _
    
    
     192:                                                      &H196C3671UL, &H6E6B06E7UL, &HFED41B76UL, &H89D32BE0UL, &H10DA7A5AUL, &H67DD4ACCUL, _
    
    
     193:                                                      &HF9B9DF6FUL, &H8EBEEFF9UL, &H17B7BE43UL, &H60B08ED5UL, &HD6D6A3E8UL, &HA1D1937EUL, _
    
    
     194:                                                      &H38D8C2C4UL, &H4FDFF252UL, &HD1BB67F1UL, &HA6BC5767UL, &H3FB506DDUL, &H48B2364BUL, _
    
    
     195:                                                      &HD80D2BDAUL, &HAF0A1B4CUL, &H36034AF6UL, &H41047A60UL, &HDF60EFC3UL, &HA867DF55UL, _
    
    
     196:                                                      &H316E8EEFUL, &H4669BE79UL, &HCB61B38CUL, &HBC66831AUL, &H256FD2A0UL, &H5268E236UL, _
    
    
     197:                                                      &HCC0C7795UL, &HBB0B4703UL, &H220216B9UL, &H5505262FUL, &HC5BA3BBEUL, &HB2BD0B28UL, _
    
    
     198:                                                      &H2BB45A92UL, &H5CB36A04UL, &HC2D7FFA7UL, &HB5D0CF31UL, &H2CD99E8BUL, &H5BDEAE1DUL, _
    
    
     199:                                                      &H9B64C2B0UL, &HEC63F226UL, &H756AA39CUL, &H26D930AUL, &H9C0906A9UL, &HEB0E363FUL, _
    
    
     200:                                                      &H72076785UL, &H5005713UL, &H95BF4A82UL, &HE2B87A14UL, &H7BB12BAEUL, &HCB61B38UL, _
    
    
     201:                                                      &H92D28E9BUL, &HE5D5BE0DUL, &H7CDCEFB7UL, &HBDBDF21UL, &H86D3D2D4UL, &HF1D4E242UL, _
    
    
     202:                                                      &H68DDB3F8UL, &H1FDA836EUL, &H81BE16CDUL, &HF6B9265BUL, &H6FB077E1UL, &H18B74777UL, _
    
    
     203:                                                      &H88085AE6UL, &HFF0F6A70UL, &H66063BCAUL, &H11010B5CUL, &H8F659EFFUL, &HF862AE69UL, _
    
    
     204:                                                      &H616BFFD3UL, &H166CCF45UL, &HA00AE278UL, &HD70DD2EEUL, &H4E048354UL, &H3903B3C2UL, _
    
    
     205:                                                      &HA7672661UL, &HD06016F7UL, &H4969474DUL, &H3E6E77DBUL, &HAED16A4AUL, &HD9D65ADCUL, _
    
    
     206:                                                      &H40DF0B66UL, &H37D83BF0UL, &HA9BCAE53UL, &HDEBB9EC5UL, &H47B2CF7FUL, &H30B5FFE9UL, _
    
    
     207:                                                      &HBDBDF21CUL, &HCABAC28AUL, &H53B39330UL, &H24B4A3A6UL, &HBAD03605UL, &HCDD70693UL, _
    
    
     208:                                                      &H54DE5729UL, &H23D967BFUL, &HB3667A2EUL, &HC4614AB8UL, &H5D681B02UL, &H2A6F2B94UL, _
    
    
     209:                                                      &HB40BBE37UL, &HC30C8EA1UL, &H5A05DF1BUL, &H2D02EF8DUL _
    
    
     210:                                                  }
    
    
     211:   
    
    
     212:  End Class

 

There it is, in all the unfriendly nastyness.  The API class I am using is calling the previous c implementation that I had before, just to verify that they match.

Just thought I would gloat a bit, it was a tedious process, and a lot of learning about .Net binary actions that I wasn't aware of, sometimes you just look at the code, but don't really understand it, I finally have a really good understanding of how this all works.

***EDIT For Credit**   
[SWGEmu.com][2]   
[CRC Explained][3]   
[Encryption Explained][4]

![][5]

[1]: http://www.renevo.com/photos/swgemudotnetscreenshots/images/169/original.aspx
[2]: http://www.swgemu.com/
[3]: http://trac2.assembla.com/swgemu/wiki/Packets/CRC
[4]: http://trac2.assembla.com/swgemu/wiki/Packets/Encryption
[5]: http://renevo.com/aggbug.aspx?PostID=1926

