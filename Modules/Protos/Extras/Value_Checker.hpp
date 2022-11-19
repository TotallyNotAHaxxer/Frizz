#pragma once 
#include <iostream>
#include <algorithm>

namespace Value_Container {
    template<class C, typename T> bool Checker(C&& c, T t) { return std::find(std::begin(c), std::end(c), t) != std::end(c); }
};