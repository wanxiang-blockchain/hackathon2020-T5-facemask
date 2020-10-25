--- sc Porject ---

 - How to Build
   - cd to 'build' directory
   - run the command export PLATONE_CDT_ROOT=<where you place cdt>
   - run the command 'cmake ..'
   - run the command 'make' 

 - After build - 
   - The built smart contract is under the 'sc' directory in the 'build' directory
   - You can then use 'ctool' to deploy contract

 - Additions to CMake should be done to CMakeLists.tx in the './src' directory and not in the top level CMakeLists.txt


 ./ctool deploy --abi /home/colin/Workspace/src/sc/build/sc/sc.abi.json --code /home/colin/Workspace/src/sc/build/sc/sc.wasm --config ../conf/ctool.json

 0x1a34170b185c9fa29287adb8ffe363d69b97e27b

 ./ctool invoke --addr 0x1a34170b185c9fa29287adb8ffe363d69b97e27b --abi /home/colin/Workspace/src/sc/build/sc/sc.abi.json  --config ../conf/ctool.json  --func "set" --param "[{"a":"ddd"},{"b":"aaa"}]" 