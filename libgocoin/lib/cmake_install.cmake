# Install script for directory: /home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin

# Set the install prefix
IF(NOT DEFINED CMAKE_INSTALL_PREFIX)
  SET(CMAKE_INSTALL_PREFIX "/usr/local")
ENDIF(NOT DEFINED CMAKE_INSTALL_PREFIX)
STRING(REGEX REPLACE "/$" "" CMAKE_INSTALL_PREFIX "${CMAKE_INSTALL_PREFIX}")

# Set the install configuration name.
IF(NOT DEFINED CMAKE_INSTALL_CONFIG_NAME)
  IF(BUILD_TYPE)
    STRING(REGEX REPLACE "^[^A-Za-z0-9_]+" ""
           CMAKE_INSTALL_CONFIG_NAME "${BUILD_TYPE}")
  ELSE(BUILD_TYPE)
    SET(CMAKE_INSTALL_CONFIG_NAME "Release")
  ENDIF(BUILD_TYPE)
  MESSAGE(STATUS "Install configuration: \"${CMAKE_INSTALL_CONFIG_NAME}\"")
ENDIF(NOT DEFINED CMAKE_INSTALL_CONFIG_NAME)

# Set the component getting installed.
IF(NOT CMAKE_INSTALL_COMPONENT)
  IF(COMPONENT)
    MESSAGE(STATUS "Install component: \"${COMPONENT}\"")
    SET(CMAKE_INSTALL_COMPONENT "${COMPONENT}")
  ELSE(COMPONENT)
    SET(CMAKE_INSTALL_COMPONENT)
  ENDIF(COMPONENT)
ENDIF(NOT CMAKE_INSTALL_COMPONENT)

# Install shared libraries without execute permission?
IF(NOT DEFINED CMAKE_INSTALL_SO_NO_EXE)
  SET(CMAKE_INSTALL_SO_NO_EXE "1")
ENDIF(NOT DEFINED CMAKE_INSTALL_SO_NO_EXE)

IF(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  IF(EXISTS "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/libzerocoin.so" AND
     NOT IS_SYMLINK "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/libzerocoin.so")
    FILE(RPATH_CHECK
         FILE "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/libzerocoin.so"
         RPATH "")
  ENDIF()
  FILE(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/lib" TYPE SHARED_LIBRARY FILES "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/lib/libzerocoin.so")
  IF(EXISTS "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/libzerocoin.so" AND
     NOT IS_SYMLINK "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/libzerocoin.so")
    IF(CMAKE_INSTALL_DO_STRIP)
      EXECUTE_PROCESS(COMMAND "/usr/bin/strip" "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/libzerocoin.so")
    ENDIF(CMAKE_INSTALL_DO_STRIP)
  ENDIF()
ENDIF(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")

IF(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  FILE(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/include/zerocoin/bitcoin_bignum" TYPE FILE FILES
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/bitcoin_bignum/allocators.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/bitcoin_bignum/bignum.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/bitcoin_bignum/clientversion.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/bitcoin_bignum/compat.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/bitcoin_bignum/hash.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/bitcoin_bignum/netbase.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/bitcoin_bignum/serialize.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/bitcoin_bignum/uint256.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/bitcoin_bignum/version.h"
    )
ENDIF(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")

IF(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  FILE(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/include/zerocoin" TYPE FILE FILES
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/Accumulator.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/AccumulatorProofOfKnowledge.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/Coin.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/CoinSpend.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/Commitment.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/ParamGeneration.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/Params.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/SerialNumberSignatureOfKnowledge.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/SpendMetaData.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/Goapi.h"
    "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/libzerocoin/Zerocoin.h"
    )
ENDIF(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")

IF(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  FILE(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/lib/pkgconfig" TYPE FILE FILES "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/lib/zerocoin.pc")
ENDIF(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")

IF(CMAKE_INSTALL_COMPONENT)
  SET(CMAKE_INSTALL_MANIFEST "install_manifest_${CMAKE_INSTALL_COMPONENT}.txt")
ELSE(CMAKE_INSTALL_COMPONENT)
  SET(CMAKE_INSTALL_MANIFEST "install_manifest.txt")
ENDIF(CMAKE_INSTALL_COMPONENT)

FILE(WRITE "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/lib/${CMAKE_INSTALL_MANIFEST}" "")
FOREACH(file ${CMAKE_INSTALL_MANIFEST_FILES})
  FILE(APPEND "/home/wkdisee/golang/work/src/github.com/wangkangda/zerochaincode/libgocoin/lib/${CMAKE_INSTALL_MANIFEST}" "${file}\n")
ENDFOREACH(file)
