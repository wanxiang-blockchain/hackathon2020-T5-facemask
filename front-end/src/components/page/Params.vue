<template>
    <div class="form-box">
        <el-form ref="form" :model="form" label-width="150px"  :label-position="'left'" style="display:block">
            <el-form-item :label="$t('i18n.gasConsumeTokenContractName')" >
                <el-input v-model="form.gasContractName" style="width:400px"></el-input>
            </el-form-item>
            <el-form-item :label="$t('i18n.txConsumeGas')" label-width="150px">
                <el-switch v-model="form.txConsumeGas"></el-switch>
            </el-form-item>
            <el-form-item :label="$t('i18n.txGasLimit')"  label-width="110px">
                <el-slider v-model="form.txGasLimit" :min="0" :max="2000000000" show-input input-size="mini" style="width:400px"></el-slider>
            </el-form-item>
            <el-form-item :label="$t('i18n.bkGasLimit')"  label-width="110px">
                <el-slider v-model="form.bkGasLimit" :min="0" :max="20000000000" show-input input-size="mini" style="width:400px"></el-slider>
            </el-form-item>
            <el-form-item :label="$t('i18n.checkDeployPermission')"  style="display:inline-block"  label-width="150px">
                <el-switch v-model="form.checkContractDeployPermission"></el-switch>
            </el-form-item>
            <el-form-item :label="$t('i18n.reviewAfterDeploy')" style="display:inline-block;margin-left:50px"  label-width="200px">
                <el-switch v-model="form.checkContractAfterDeploy"></el-switch>
            </el-form-item>
            <el-form-item :label="$t('i18n.isEmptyBlock')"  style="display:inline-block"   label-width="150px">
                <el-switch v-model="form.allowEmpty"></el-switch>
            </el-form-item>
            <el-form-item :label="$t('i18n.firewall')" style="display:inline-block;margin-left:50px"  label-width="150px">
                <el-switch v-model="form.fw" disabled></el-switch>
            </el-form-item>
            <el-form-item style="margin-top:15px" label-width="230px">
                <el-button type="primary" @click="onSubmit">{{$t('i18n.submitChanges')}}</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
import {dataUrl,apiServerUrl} from '../../config'
export default {
    data() {
        return {
            form: {
                gasContractName: '',
                allowEmpty:1==0?false:true,
                txGasLimit:0,
                bkGasLimit:0,
                checkContractDeployPermission:1==0?false:true,
                checkContractAfterDeploy:1==0?false:true,
                txConsumeGas:1==0?false:true,
                fw:false
            },
        };
    },
    methods: {
        getData(){
            this.axios.get(apiServerUrl + '/superadmin/systemconfig/getsystemconfig')
            .then(res => {
                if (res && res.data) {
                   let resData = JSON.parse(res.data.data)
                   console.log(resData)
                   this.form.gasContractName = resData.gasContractName.replace('""','');
                   this.form.txGasLimit = parseInt(resData.txGasLimit);
                   this.form.bkGasLimit = parseInt(resData.blockGasLimit);

                   this.form.allowEmpty = resData.isProduceEmptyBlock == "1";
                   this.form.checkContractDeployPermission = resData.isCheckDeployPermission == "1";
                   this.form.checkContractAfterDeploy = resData.isApproveDeployedContract == "1";
                   this.form.txConsumeGas = resData.isUseGas == "1";
                }
            }).catch(res => {
                console.log(res)
            })
        },
        onSubmit() {
            this.axios.get(apiServerUrl + '/superadmin/systemconfig/getsystemconfig')
            .then(res => {
                if (res && res.data) {
                    let resData = JSON.parse(res.data.data)
                    let setparams = {}

                    if (this.form.gasContractName != resData.gasContractName.replace('""','')){
                        setparams.GasContractName = this.form.gasContractName;
                    }
                    if (this.form.txGasLimit != parseInt(resData.txGasLimit)){
                        setparams.TxGasLimit = this.form.txGasLimit;
                    }
                    if (this.form.bkGasLimit != parseInt(resData.blockGasLimit)){
                        setparams.BlockGasLimit = this.form.bkGasLimit;
                    }
                    if (this.form.allowEmpty != resData.isProduceEmptyBlock == "1"){
                        setparams.IsProduceEmptyBlock = this.form.allowEmpty?1:0;
                    }
                    if (this.form.checkContractDeployPermission != resData.isCheckDeployPermission == "1"){
                        setparams.IsCheckDeployPermission = this.form.checkContractDeployPermission?1:0;
                    }
                    if (this.form.checkContractAfterDeploy != resData.isApproveDeployedContract == "1"){
                        setparams.IsApproveDeployedContract = this.form.checkContractAfterDeploy?1:0;
                    }
                    if (this.form.txConsumeGas != resData.isUseGas == "1"){
                        setparams.IsUseGas = this.form.txConsumeGas?1:0;
                    }

                    if (this.form.txGasLimit >= this.form.bkGasLimit) {
                        this.$message.error("txGasLimit must not bigger than bkGasLimit");
                    }

                    this.axios.get(apiServerUrl + '/superadmin/systemconfig/setsystemconfig',{
                        params:setparams
                    }).then(res => {
                        console.log(res)
                        this.$message.success(this.$t('i18n.submitSuccess'));
                    }).catch(res =>{
                        this.$message.error(this.$t('i18n.submitFailed'));
                    })
                }
            }).catch(res => {
                console.log(res)
            })
        }
    },
    mounted:function() {
        this.getData();
        setTimeout(this.getData,1000)
    },
};
</script>
<style scoped>
.form-box >>> .el-input-number--mini .el-input__inner{
    padding-left: 10px!important;
    padding-right: 10px;
}
.form-box >>> .el-input-number__decrease{
    width:23px;
}
.form-box >>> .el-input-number__increase{
    width:23px;
}
</style>