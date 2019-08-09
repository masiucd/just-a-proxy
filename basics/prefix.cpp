#include <iostream>
#include <math.h>
#include <cstdlib>

int main()
{
  int data;
  data = 5;

  std::cout << "data has the value of = " << data * data << std::endl;

  data++;
  std::cout << data << std::endl;
  ++data;
  std::cout << data << std::endl;
  data--;
  std::cout << data << std::endl;

  return 0;
}
