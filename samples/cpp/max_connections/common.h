#pragma once
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <arpa/inet.h>
#include <iostream>

#define ERROR(args ...)  fprintf(stderr, args)

using namespace std;

char* ADDR = "0.0.0.0";
unsigned short int PORT = 8088;



