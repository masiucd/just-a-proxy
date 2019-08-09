#include <iostream>
using namespace std;

int main()
{
  int age;

  cout << "How old are you ?" << endl;
  cin >> age;

  if (age < 18)
  {
    cout << "Sorry but you can't drink" << endl;
  }
  else if (age == 18)
  {
    cout << "Welcome to the party" << endl;
  }
  else if (age > 18)
  {
    cout << "You have done this for a long time now!!!" << endl;
  }
  system("PAUSE");
  return EXIT_SUCCESS;
}