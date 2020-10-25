include(ExternalProject)
# if no cdt root is given use default path
if (PLATONE_CDT_ROOT STREQUAL "" OR NOT PLATONE_CDT_ROOT)
  find_package(platone.cdt)
endif()

ExternalProject_Add(
  sc_project
  SOURCE_DIR ${CMAKE_SOURCE_DIR}/src
  BINARY_DIR ${CMAKE_BINARY_DIR}/sc
  CMAKE_ARGS -DCMAKE_TOOLCHAIN_FILE=${PLATONE_CDT_ROOT}/lib/cmake/platone.cdt/PlatoneWasmToolchain.cmake -DPLATONE_CDT_ROOT=${PLATONE_CDT_ROOT}
  UPDATE_COMMAND ""
  PATCH_COMMAND ""
  TEST_COMMAND ""
  INSTALL_COMMAND ""
  BUILD_ALWAYS 1
)