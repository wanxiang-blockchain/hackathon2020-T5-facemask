<template>
<div>
    <div slot="header">
        <span>{{$t('i18n.userList')}}</span>
        <el-button style="float: right; padding: 3px 0" type="text" @click="openDialog">{{$t('i18n.add')}}</el-button>
    </div>
    <el-scrollbar style="height:100%;">
    <el-table :data="userList" :max-height="height">
        <el-table-column :label="`${$t('i18n.account')}`" :width="220" show-overflow-tooltip>
            <template slot-scope="scope">
                 <el-button type="text" @click="onRowClick(scope.row)">{{scope.row.name}}</el-button>
            </template>
        </el-table-column>
        <el-table-column prop="role" :label="`${$t('i18n.role')}`" show-overflow-tooltip ></el-table-column>
        <el-table-column prop="registertime" :label="`${$t('i18n.createDate')}`" show-overflow-tooltip ></el-table-column>
    </el-table>
    </el-scrollbar>

    <el-dialog :title="$t('i18n.user')" :visible.sync="dialogFormVisible">
        <el-form :model="form" label-width="100px" :rules="rules" ref="adduser" :label-position="'left'">
            <el-form-item :label="$t('i18n.account')" style="width:350px" prop="username">
                <el-input v-model="form.username" placeholder="" :disabled="inputDisabled"></el-input>
            </el-form-item>
            <el-form-item :label="$t('i18n.password')" style="width:350px" prop="password" v-if="!inputDisabled">
                <el-input v-model="form.password" placeholder=""></el-input>
            </el-form-item>
            <el-form-item :label="$t('i18n.password')" style="width:350px" prop="newpassword" v-if="inputDisabled">
                <el-input v-model="form.newpassword" placeholder=""></el-input>
            </el-form-item>
            <el-form-item :label="$t('i18n.role')">
                <el-select v-model="form.role" placeholder="">
                    <el-option key="USER" :label="$t('i18n.user')" value="USER"></el-option>
                    <el-option key="SUPERADMIN" :label="$t('i18n.admin')" value="SUPERADMIN"></el-option>
                </el-select>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="cancel">{{$t('i18n.cancel')}}</el-button>
            <el-button type="primary" @click="submitForm">{{$t('i18n.confirm')}}</el-button>
        </div>
    </el-dialog>
</div>
</template>

<script>
import {dataUrl,apiServerUrl} from '../../config'
export default {
    data: function() {
        return {
            userList:[],
            dialogFormVisible:false,
            inputDisabled : false,
            form:{
                username:"",
                password:"",
                newpassword:"",
                role:"USER"
            }
        }
    },
    methods: {
        getData(){
            this.axios.get(apiServerUrl + '/superadmin/getusers',{
                params: {
                            PageIndex:1,
                            PageSize:10000
                        }
            })
            .then(res => {
                console.log(res.data);
                this.userList = res.data.data;
            }).catch(res => {
                console.log(res);
            });
        },
        openDialog(){
            this.resetForm()
            this.dialogFormVisible = true
            this.inputDisabled = false
        },
        onRowClick(row){
            this.form = {
                username:row.name,
                password:"",
                role:row.role
            }
            this.dialogFormVisible = true
            this.inputDisabled = true
        },
        submitForm(){
            this.$refs.adduser.validate(valid => {
                if (valid) {
                    let url = apiServerUrl + '/superadmin/adduser';
                    if (this.inputDisabled){
                        url = apiServerUrl + '/superadmin/update'
                    }
                    this.axios.get(url,{
                        params: {
                            Name:this.form.username,
                            Password:this.form.password,
                            NewPassword:this.form.newpassword,
                            Role:this.form.role,
                            Passphrase:'0'
                        }
                    }).then(res => {
                    if (res && res.data) {
                        console.log(res.data)
                        this.$message.success(this.$t('i18n.submitSuccess'));
                        this.getData()
                    }
                    }).catch(res => {
                        this.$message.error(this.$t('i18n.submitFailed'));
                    })
                } else {
                    this.$message.error(this.$t('i18n.enterNameAndPassword'));
                }
            });

            this.resetForm()
            this.dialogFormVisible = false
        },
        cancel(){
            this.resetForm()
            this.dialogFormVisible = false
        },
        resetForm(){
            this.form = {
                username:"",
                password:"",
                newpassword:"",
                role:"USER"
            }
        }
    },
    mounted:function() {
        this.getData()
        setInterval(this.getData(),2000)
    },
    computed:{
        rules() {
            return {
                username: [{ required: true, message: this.$t('i18n.enterName'), trigger: 'blur' }],
                password: [{ required: true, message: this.$t('i18n.enterPassword'), trigger: 'blur' }],
            }
        }
    },
    props:{
        height:{
            default:270
        }
    }
};
</script>

<style scoped>
.ms-content {
    padding: 30px 30px 10px 30px;
}

.btn{
    width:100%
}

a{
    color: #3498db;
}

.button-custom{
    padding:0px;
    border: 0px;
}
</style>

