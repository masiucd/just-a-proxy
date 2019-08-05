#include <iostream>
#include "Tweeter.h"
#include "Person.h"

using std::cout;
using std::endl;
using std::string;

int main()
{
  Person p1("Boris", "Arshavin", 16);
  {
    Tweeter t1("Hello ", " foo ", 456, "@bar");
    std::string name2 = ti.getName();
  }
  cout << "after innermost block " << endl;
  string name = p1.getName();

  return 0;
}