#include <iostream>
using namespace std;
// using std::cout;
// using std::endl;
// using std::string;

class Human
{
public:
  string name;
  int age;

  void greet()
  {
    cout << "my name is " << name << " and I am " << age << " years old" << endl;
  }

  void eat(string food)
  {
    cout << name << " is eating " << food << endl;
  }

  void eat(string f, int amount)
  {
    cout << name << " is eating " << f << " with a amount of: " << amount << endl;
  }
};

int main()
{
  Human exhibitA;
  Human exhibitB;

  exhibitA.name = "Marcello";
  exhibitA.age = 22;

  exhibitB.name = "Logan";
  exhibitB.age = 25;

  cout << "Exhibit A's Name is: " << exhibitA.name << endl;
  cout << "Exhibit B's age is: " << exhibitA.age << endl;

  cout << "Exhibit B's Name is: " << exhibitB.name << endl;
  cout << "Exhibit B's age is: " << exhibitB.age << endl;

  exhibitA.greet();
  exhibitB.eat("Pizza");
  exhibitA.eat("Bananas", 5);

  return 0;
}