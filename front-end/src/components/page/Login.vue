<template>
    <div class="login-wrap">
        <div class="ms-login">
            <div class="ms-title">{{$t('i18n.system_name')}}</div>
            <el-form :model="param" :rules="rules" ref="login" label-width="0px" class="ms-content">
                <el-form-item prop="username">
                    <el-input v-model="param.username" placeholder="username">
                        <el-button slot="prepend" icon="el-icon-lx-people"></el-button>
                    </el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input
                        type="password"
                        placeholder="password"
                        v-model="param.password"
                        @keyup.enter.native="submitForm()"
                    >
                        <el-button slot="prepend" icon="el-icon-lx-lock"></el-button>
                    </el-input>
                </el-form-item>
                <div class="login-btn">
                    <el-button type="primary" @click="submitForm()">{{$t('i18n.login')}}</el-button>
                </div>
            </el-form>
        </div>
    </div>
</template>

<script>
import {apiServerUrl} from '../../config'
export default {
    data: function() {
        return {
            param: {
                username: 'admin',
                password: 'admin',
            }
        };
    },
    methods: {          
        submitForm() {
            this.$refs.login.validate(valid => {
                // if (process.env.NODE_ENV == 'development'){
                //     this.$message.success(this.$t('i18n.loginSuccess'));
                //         localStorage.setItem('ms_username', 'root');
                //         localStorage.setItem('ms_userrole', 'superadmin');
                //         this.$router.push('/');
                //         return;
                // }

                if (valid) {
                    this.axios.get(apiServerUrl + '/login',{
                        params: {
                            Name:this.param.username,
                            Password:this.param.password
                        }
                    }).then(res => {
                    if (res && res.data) {
                        console.log(res.data)
                        this.$message.success(this.$t('i18n.loginSuccess'));
                        localStorage.setItem('ms_username', this.param.username);
                        localStorage.setItem('ms_last_login_time', res.data.data[0]);
                        localStorage.setItem('ms_last_login_ip', res.data.data[1]);
                        localStorage.setItem('ms_userrole', res.data.data[2].toLowerCase());
                        this.$router.push('/');
                    }
                    }).catch(res => {
                        this.$message.error(this.$t('i18n.loginFailed'));
                    })
                } else {
                    this.$message.error(this.$t('i18n.enterNameAndPassword'));
                }
            });
        }
    },
    computed:{
        rules() {
            return {
                username: [{ required: true, message: this.$t('i18n.enterName'), trigger: 'blur' }],
                password: [{ required: true, message: this.$t('i18n.enterPassword'), trigger: 'blur' }],
            }
        }
    }
};
</script>

<style scoped>
.login-wrap {
    position: relative;
    width: 100%;
    height: 100%;
    background-image: url(../../assets/img/login-bg.jpg);
    background-size: 100%;
}
.ms-title {
    width: 100%;
    line-height: 50px;
    text-align: center;
    font-size: 20px;
    color: rgb(10, 127, 223);
    border-bottom: 1px solid #ddd;
}
.ms-login {
    position: absolute;
    left: 50%;
    top: 50%;
    width: 350px;
    margin: -190px 0 0 -175px;
    border-radius: 5px;
    background: rgba(255, 255, 255, 0.3);
    overflow: hidden;
}
.ms-content {
    padding: 30px 30px;
}
.login-btn {
    text-align: center;
}
.login-btn button {
    width: 100%;
    height: 36px;
    margin-bottom: 10px;
}
.login-tips {
    font-size: 12px;
    line-height: 30px;
    color: #fff;
}
</style>