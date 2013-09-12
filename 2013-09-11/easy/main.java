/**
 * Class to run da easy challonge
 * @author Greg "da Beard" Cremins
 * @version 9-11-2013
 */

public class main
{
  /**
   * Method to test if there is a "+(char)+" everwhere there is a char
   * @param String test the string to be tested
   * @return the result if the pattern was repeated
   */
  public static boolean simpleSymbols(String test)
  {
    //flag variable to be carried throughout the program
   boolean flag = true;
   if(test.length() > 0)
   {
   for(int i = 0; i < test.length(); i++)
   {
    if((test.charAt(i) != '=') && (test.charAt(i) != '+'))
    {
      //if the character at the end or the beginning is not a + or an =, return false
      if((i == 0) || (i == test.length() - 1)  )
      {
        flag = false; 
        return flag;
      }
      //otherwise, return the test of the pattern
      else 
      {
        flag = (test.charAt(i - 1) == '+' &&  test.charAt(i + 1) == '+');
      }
    }
   }
   return flag;
   }
   else
   {
     return false;
   }
  }
  
  public static void main (String [] args)
  {
  String a = "++d+===+c++==a";
  System.out.println(simpleSymbols(a));
  String b = "+d+";
  System.out.println(simpleSymbols(b));
  String c = "+d+======+C+=====+r+=====";
  System.out.println(simpleSymbols(c));
  
  }
  
}