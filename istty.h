#ifdef _WIN32
#if       _WIN32_WINNT < 0x0500
  #undef  _WIN32_WINNT
  #define _WIN32_WINNT   0x0500
#endif
#include <windows.h>
#include <stdbool.h>
#include "Wincon.h" 

bool isTTY() {
  HWND consoleWnd = GetConsoleWindow();
  DWORD dwProcessId;
  GetWindowThreadProcessId(consoleWnd, &dwProcessId);
  if (GetCurrentProcessId()==dwProcessId) {
    return false;
  }
  return true;
}

void closeTTY() {
  ShowWindow(GetConsoleWindow(), SW_HIDE);
}
#elif defined(__APPLE__) || defined(__linux) || defined(__unix) || defined(__posix)
#include <stdio.h>
#include <unistd.h>
#include <stdbool.h>

bool isTTY() {
  if (isatty(fileno(stdin))) {
    return true;
  } 
  return false;
}

void closeTTY() {}
#endif