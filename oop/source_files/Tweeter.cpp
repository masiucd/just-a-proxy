#include "Tweeter.h"
#include <iostream>

Tweeter::Tweeter(std::string first,
                 std::string last,
                 int arbitrary,
                 std::string handle)
    : Person(first, last, arbitrary), twitterHandle(handle)
{
  std::cout << "constructing tweeter " << std::endl;
}

Tweeter::~Tweeter()
{
  std::cout << "destructing Tweeter " << twitterHandle << std::endl;
}