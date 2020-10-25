<template>
<div>
    <el-form :model="param" :rules="rules" ref="search" label-width="0px" class="ms-content" @submit.native.prevent>
        <el-form-item prop="info">
            <el-input v-model="param.info" placeholder="Search by contract address" @keyup.enter.native="submitForm()">
                <el-button slot="prepend" icon="el-icon-lx-search"></el-button>
            </el-input>
        </el-form-item>
    </el-form>
    <div style="padding:0px 30px 10px 30px">
        <el-card shadow="hover" >
            <el-button type="primary" round style="margin-bottom:10px" @click="dialogFormVisible = true">{{$t('i18n.deploy')}}</el-button>

            <el-dialog :title="$t('i18n.deploy')" :visible.sync="dialogFormVisible">
            <el-form :model="form">
                <el-form-item :label="$t('i18n.contractFile')">
                    <div style="display:inline-block">
                         <el-upload
                            ref="upload"
                            drag
                            action=""
                            multiple
                            :http-request="httpRequest"
                            :auto-upload="false">
                                <i class="el-icon-upload"></i>
                                <div class="el-upload__text">{{$t('i18n.dragFile')}}<em>{{$t('i18n.clickUpload')}}</em></div>
                                <div class="el-upload__tip" slot="tip">{{$t('i18n.twoFile')}}</div>
                        </el-upload>
                    </div>
                </el-form-item>
                 <el-form-item :label="$t('i18n.isSetCNS')">
                    <el-switch v-model="form.setCNS"></el-switch>
                    <el-input v-model="form.CNS" placeholder="CNS Name" v-if="form.setCNS"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogFormVisible = false">{{$t('i18n.cancel')}}</el-button>
                <el-button type="primary" @click="submitDeploy">{{$t('i18n.confirm')}}</el-button>
            </div>
            </el-dialog>

            <el-table :data="contracts">
                <el-table-column :label="`${$t('i18n.scName')}`" :max-width="250" show-overflow-tooltip>
                    <template slot-scope="scope"><router-link :to = "'/contract/' + scope.row.address">{{'(' + scope.row.name + ')' + scope.row.address}}</router-link></template>
                </el-table-column>
                <el-table-column :label="`Creator`" show-overflow-tooltip  :max-width="250">
                    <template slot-scope="scope"><router-link :to = "'/address/' + scope.row.creator + '/txs'">{{scope.row.creator}}</router-link></template>
                </el-table-column>
                <el-table-column :label="`Transaction`"  show-overflow-tooltip :max-width="250">
                    <template slot-scope="scope"><router-link :to = "'/tx/' + scope.row.tx_hash">{{scope.row.tx_hash}}</router-link></template>
                </el-table-column>
                <el-table-column :label="`${$t('i18n.time')}`"  :width="150">
                    <template slot-scope="scope">
                        <span>{{(new Date(scope.row.timestamp *1000)).toLocaleString()}}</span>
                    </template>  
                </el-table-column>
                </el-table>
            <el-pagination  
                @current-change="handleCurrentChange" 
                @size-change="handleSizeChange"
                :current-page="currentPage" 
                :page-size="PageSize" layout="total, sizes, prev, pager, next, jumper" 
                :total="totalCount">
              </el-pagination>
        </el-card>
    </div>
</div>
    
</template>

<script>
import {dataUrl,apiServerUrl} from '../../config'
export default {
    data: function() {
        return {
            param: {
                info: '',
            },
            contracts:[],
            currentPage:1,
            PageSize:10,
            totalCount:100,
            form: {
                setCNS:false,
                CNS:""
            },
            file:[],
            dialogFormVisible: false
        }
    },
    methods: {
        submitForm() {
            this.$refs.search.validate(valid => {
                if (valid) {
                    this.axios.get(dataUrl + '/contract/' + this.param.info)
                    .then(res => {
                        if(res.data && res.data.address){
                            this.$router.push('/contract/' + res.data.address);
                        }
                    }).catch(res => {
                        return false;
                    });
                } else {
                    this.$message.error(this.$t("i18n.searchInfo"));
                    return false;
                }
            });
        },
        httpRequest(param){
            this.file.push(param.file)
        },

        submitDeploy(){
            this.$refs.upload.submit();

            var updata = new FormData();
            this.file.forEach(function(file){
                updata.append('file',file,file.name)
            })
            updata.append('body',JSON.stringify(this.form))
            console.log(updata.getAll('file'))
            console.log(updata.get('body'))

            this.axios.post(apiServerUrl + '/user/deploycontract',updata,{
                headers:{'Content-Type': 'multipart/form-data' }
            }).then(res => {
                console.log(res);

                //set cns
                if (this.form.setCNS &&this.form.CNS != ""){

                    let data = res.data.data.replace(/\\n/g,'').replace(/\\t/g,'')
                    let dataJson = JSON.parse(JSON.parse(data))
                    console.log(dataJson)
                    this.axios.get(apiServerUrl + '/user/cns/setcnsnameforaddress',{
                        params: {
                            Name:this.form.CNS,
                            Version:'0.0.0.1',
                            Address:dataJson.contractAddress
                        }
                    })
                }
                this.$message.success(this.$t('i18n.submitSuccess'));
                window.location.reload();
            }).catch(res => {
                console.log(res);
            });
            this.dialogFormVisible = false
        },
        onRowClick (row, column, event){
            alert(row.miner)
        },
        getData(){
            this.axios.get(dataUrl + '/contracts',{
                    params: {
                        page_index: this.currentPage,
                        page_size: this.PageSize
                    }
                }).then(res => {
                this.contracts = res.data.items;
                this.totalCount = res.data.total;
            }).catch(res => {
                console.log(res);
            });
        },
        handleCurrentChange(val){
            this.currentPage = val;
            this.getData();
        },
        handleSizeChange(val){
            this.PageSize = val;
             this.getData();
        },
    },
    mounted:function() {
        this.getData()
        setInterval(this.getData(),2000)
    },    
    computed:{
        rules()
        {
            return {
                info: [{ required: true, message: this.$t("i18n.searchInfo"), trigger: 'blur' }]
            }    
        },
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
</style>

