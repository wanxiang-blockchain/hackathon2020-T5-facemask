<template>
<div>
    <el-form :model="param" :rules="rules" ref="search" label-width="0px" class="ms-content" @submit.native.prevent>
        <el-form-item prop="info">
            <el-input v-model="param.info" placeholder="Search by Transaction ID" @keyup.enter.native="submitForm()">
                <el-button slot="prepend" icon="el-icon-lx-search"></el-button>
            </el-input>
        </el-form-item>
    </el-form>
    <div style="padding:0px 30px 10px 30px">
        <el-card shadow="hover" >
            <el-table :data="txList">
                <el-table-column :label="`${$t('i18n.txh')}`" :width="140" show-overflow-tooltip>
                    <template slot-scope="scope"><router-link :to = "'/tx/' + scope.row.tx_hash">{{scope.row.tx_hash}}</router-link></template>
                </el-table-column>
                <el-table-column :label="`${$t('i18n.bkh')}`" :width="80" show-overflow-tooltip>
                    <template slot-scope="scope"><router-link :to = "'/block/' + scope.row.block_height">{{scope.row.block_height}}</router-link></template>
                </el-table-column>
                <el-table-column :label="'From'" :max-width="260" show-overflow-tooltip >
                    <template slot-scope="scope"><router-link :to = "'/address/' + scope.row.from + '/txs'">{{scope.row.from}}</router-link></template>
                </el-table-column>
                <el-table-column :label="'To'"  :max-width="260" show-overflow-tooltip>
                    <template slot-scope="scope"><router-link :to = "'/address/' + scope.row.to + '/txs'">{{scope.row.to}}</router-link></template>
                </el-table-column>
                <el-table-column :label="`${$t('i18n.gasUsed')}`">
                <template slot-scope="scope">
                    <span>{{scope.row.gasUsed}}</span>
                    <span>{{'('+ (scope.row.gas_limit == 0?0:(scope.row.gas_used/scope.row.gas_limit*100)).toFixed(4) + '%)'}}</span>
                    <el-progress :show-text="false" :stroke-width="5" :percentage="scope.row.gas_limit == 0?0:(scope.row.gas_used/scope.row.gas_limit*100)" ></el-progress>
                </template>    
            </el-table-column>
            <el-table-column :label="`${$t('i18n.time')}`" :width="160"> 
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
import {dataUrl} from '../../config'
export default {
    data: function() {
        return {
            param: {
                info: '',
            },
            txList:[],
            currentPage:1,
            PageSize:10,
            totalCount:100
        }
    },
    props:['blockNum','addr'],
    methods: {
        submitForm() {
            this.$refs.search.validate(valid => {
                if (valid) {
                    this.axios.get(dataUrl + '/tx/' + this.param.info)
                    .then(res => {
                        if(res.data && res.data.tx_hash){
                            this.$router.push('/tx/' + res.data.tx_hash);
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
        getData(){
            var url = dataUrl + '/txs'
            if (this.blockNum){
                url = dataUrl + "/block/" + this.blockNum + "/txs"
            }
            else if(this.addr){
                url = dataUrl + "/address/from/" + this.addr + "/txs"
            }
            console.log(url)
            this.axios.get(url,{
                    params: {
                        page_index: this.currentPage,
                        page_size: this.PageSize
                    }
                }).then(res => {
                    console.log(res.data.items)
                this.txList = res.data.items;
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
        }
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
    },
    watch: {
        '$route' (to, from) {
            this.getData()
        }
    },
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

