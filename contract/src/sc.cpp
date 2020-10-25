//auto create contract
#include <stdlib.h>
#include <string.h>
#include <string>
#include <bcwasm/bcwasm.hpp>

namespace demo { 
class sc : public bcwasm::Contract{
    public:
    sc(){}

    /// 实现父类: bcwasm::Contract 的虚函数
    /// 该函数在合约首次发布时执行，仅调用一次
    void init()
    {
        bcwasm::println("init success...");
    }
    /// 定义Event.
    /// BCWASM_EVENT(eventName, arguments...)
    BCWASM_EVENT(set, const char *)

    public:
    void set(const char *msg)
    {
        // 定义状态变量
        bcwasm::setState("NAME_KEY", std::string(msg));
        bcwasm::println("setName success...");

        // 事件返回
      BCWASM_EMIT_EVENT(set, "std::string(msg)");
      bcwasm::println("emit event success...");
    }
    const char* get() const
    {
        std::string value;
        bcwasm::getState("NAME_KEY", value);
        // 读取合约数据并返回
        bcwasm::println(value);
        return value.c_str();
    }
};
}
// 此处定义的函数会生成ABI文件供外部调用
BCWASM_ABI(demo::sc, set)
BCWASM_ABI(demo::sc, get)
BCWASM_ABI(demo::sc, init)
