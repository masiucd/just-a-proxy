#include <iostream>
using namespace std;

int main()
{
  const int numOfNames = 4;
  string names[numOfNames] = {"hans", "tim", "boris", "karol"};
  int i;
  for (i = 0; i < 10; i++)
  {
    cout << " i = " << i << endl;
  }

  for (i = 0; i < numOfNames; i++)
  {
    cout << " names = " << names << endl;
  }
  {
    /* code */
  }

  return 0;
}