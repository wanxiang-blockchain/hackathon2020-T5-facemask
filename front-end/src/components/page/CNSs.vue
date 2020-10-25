<template>
<div>
    <div slot="header">
        <span>{{$t('i18n.cnsList')}}</span>
        <el-button style="float: right; padding: 3px 0" type="text" @click="openDialog">{{$t('i18n.add')}}</el-button>
    </div>
        <el-scrollbar style="height:100%;">
        <el-table :data="cnsList" :max-height="height" style="width:100%">
        <el-table-column :label="`${$t('i18n.scName')}`" :width="130">
            <template slot-scope="scope">
                 <el-button type="text" @click="onRowClick(scope.row)">{{scope.row.name}}</el-button>
            </template>
        </el-table-column>
        <el-table-column prop="version" :label="`${$t('i18n.currVersion')}`"  width="130"></el-table-column>
        <el-table-column prop="address" :label="`${$t('i18n.currAddr')}`" ></el-table-column>
        </el-table>
        </el-scrollbar>


        <el-dialog :title="'CNS'" :visible.sync="dialogFormVisible">
        <el-form :model="form" label-width="150px"  :label-position="'left'">
            <el-form-item :label="$t('i18n.cnsContractName') + ':'" style="width:350px">
                <el-input v-model="form.name" :placeholder="$t('i18n.pinput')" :disabled="isEdit"></el-input>
            </el-form-item>
            <el-form-item :label="$t('i18n.version') + ':'" style="width:350px" v-if="!isEdit">
                <el-input v-model="form.version" :placeholder="$t('i18n.version') +' 0.0.0.1'"></el-input>
            </el-form-item>
            <el-form-item :label="$t('i18n.version') + ':'"  v-if="isEdit">
                <div v-for="info in form.infos" v-bind:key="info.version">
                    <el-radio v-model="form.version" :label="info.version" @change="onRadioChange">{{info.version}}</el-radio>
                </div>
            </el-form-item>
            <el-form-item :label="$t('i18n.scAddr')+ ':'" >
                 <el-input v-model="form.address" :placeholder="$t('i18n.scAddr')" :disabled="isEdit"></el-input>
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
    data() {
        return {
            cnsList: [],
            dialogFormVisible : false,
            form:{
                name:"",
                version:"",
                address:"",
                infos:[]
            },
            isEdit:false
        }
    },
    methods:{
        getData(){
            this.axios.get(dataUrl + '/cns',{
                    params: {
                        page_index: 1,
                        page_size: 10000
                    }
                }).then(res => {
                if (res && res.data) {
                    this.cnsList = res.data.items;
                }
            }).catch(res => {
                console.log(res)
            })
        },
        openDialog(){
            this.resetForm()
            this.isEdit = false;
            this.dialogFormVisible = true
        },
        onRowClick(row){
            this.form = {
                name:row.name,
                version:row.version,
                address:row.address,
                infos:row.infos
            }
            this.isEdit = true
            this.dialogFormVisible = true
        },
        submitForm(){
            let url = '/superadmin/cns/setcnsnameforaddress'
            if (this.isEdit){
                url = '/superadmin/cns/cnsredirect'
            }
            this.axios.get(apiServerUrl + url,{
                params: {
                    Name:this.form.name,
                    Version:this.form.version,
                    Address:this.form.address
                }
            }).then(res => {
                this.$message({
                    type: 'success',
                    message: this.$t('i18n.submitSuccess')
                });
            }).catch(res => {
                this.$message({
                    type: 'error',
                    message: this.$t('i18n.submitFailed')
                });
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
                name:"",
                version:"",
                address:"",
                infos:[]
            }
        },
        onRadioChange(val){
            for  (let i = 0;i < this.form.infos.length ;i++){
                let info = this.form.infos[i];
                if (info.version == val){
                    this.form.address = info.address;
                }
            }
        }
    },
    mounted:function() {
        this.getData()
        setInterval(this.getData,5000)
    },
    props:{
        height:{
            default:270
        }
    }
}
</script>

<style scoped>

</style>