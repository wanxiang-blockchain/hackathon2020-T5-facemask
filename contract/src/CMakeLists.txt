project(sc)

if (PLATONE_CDT_ROOT STREQUAL "" OR NOT PLATONE_CDT_ROOT)
  find_package(platone.cdt)
else()
  list(APPEND CMAKE_MODULE_PATH ${PLATONE_CDT_ROOT}/lib/cmake/platone.cdt)
  include(PlatoneWasmToolchain)
  include(PlatoneCDTMacros)
endif()

add_contract(sc sc sc.cpp)
target_include_directories(sc PUBLIC ${CMAKE_SOURCE_DIR}/../include)
