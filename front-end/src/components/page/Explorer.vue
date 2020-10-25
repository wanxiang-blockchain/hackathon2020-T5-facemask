<template>
<div class = "explorer">
    <el-row :gutter="20">
        <el-col :span="12">
                <el-card shadow="hover">
                    <div slot="header" style="margin-bottom:10px">
                        <span>{{$t('i18n.bkList')}}</span>
                    </div>
                    <el-table :data="bkList" style="width:100%;"  max-height="550" :show-header="false">
                    <el-table-column :width="220">
                        <template slot-scope="scope">
                            <div class="bk_icon"><span>BK</span></div>
                            <div class="bk_icon_right">
                                <router-link :to = "'/block/' + scope.row.height">{{scope.row.height}}</router-link>
                                <div>{{new Date(scope.row.timestamp *1000).toLocaleString()}}</div>
                            </div>
                        </template>
                    </el-table-column>
                    <el-table-column>
                        <template slot-scope="scope">
                            <div style="display:block">
                                <span style="font-weight:bold;">Miner </span>
                                <router-link :to = "'/address/' + scope.row.proposer + '/txs'">{{scope.row.proposer}}</router-link>
                            </div>
                            <div style="display:block">
                                <router-link :to = "'/block/' + scope.row.height + '/txs'">{{scope.row.tx_amount}}</router-link>
                                 <span style="">  txns in this block </span>
                            </div>
                        </template>
                    </el-table-column>
                    </el-table>
                    <div>
                        <el-button type="primary" class="btn" @click="onClickBk">{{$t('i18n.viewAllBks')}}</el-button>
                    </div>
                </el-card>
        </el-col>
        <el-col :span="12">
            <el-card shadow="hover">
                <div slot="header" style="margin-bottom:10px">
                        <span>{{$t('i18n.txList')}}</span>
                </div>
                <el-table :data="txList" max-height="550" :show-header="false">
                <el-table-column :width="330" show-overflow-tooltip>
                    <template slot-scope="scope">
                            <div class="bk_icon tx_icon"><span>TX</span></div>
                            <div class="bk_icon_right">
                                <router-link :to = "'/tx/' + scope.row.tx_hash" class="txhash">{{scope.row.tx_hash}}</router-link>
                                <div>{{(new Date(scope.row.timestamp *1000)).toLocaleString()}}</div>
                            </div>
                        </template>
                </el-table-column>
                <el-table-column  show-overflow-tooltip >
                    <template slot-scope="scope">
                            <div class="txhash">
                                <span style="font-weight:bold;">From: </span>
                                <router-link :to = "'/address/' + scope.row.from + '/txs'">{{scope.row.from}}</router-link>
                            </div>
                            <div class="txhash">
                                <span style="font-weight:bold;">To: </span>
                                <router-link :to = "'/address/' + scope.row.to + '/txs'">{{scope.row.to}}</router-link>
                            </div>
                        </template>
                </el-table-column>
                </el-table>
                <div>
                    <el-button type="primary" class="btn"  @click="onClickTx">{{$t('i18n.viewAllTxs')}}</el-button>
                </div>
            </el-card>
        </el-col>
    </el-row>
</div>
</template>

<script>
import {dataUrl} from '../../config'
export default {
    data: function() {
        return {
            bkList:[],
            txList:[]
        }
    },
    methods: {
        onClickBk(){
            this.$router.push('/blocks');
        },
        onClickTx(){
            this.$router.push('/txs');
        },
        getData(){
            this.axios.get(dataUrl + '/blocks',{
                    params: {
                        page_index: 1,
                        page_size:7
                    }
                }).then(res => {
                this.bkList = res.data.items;
            }).catch(res => {
                console.log(res);
            });

            this.axios.get(dataUrl + '/txs',{
                    params: {
                        page_index: 1,
                        page_size:7
                    }
                }).then(res => {
                this.txList = res.data.items;
            }).catch(res => {
                console.log(res);
            });
        }
    },
    mounted:function() {
        this.getData()
        setInterval(this.getData,2000)
    }
};
</script>

<style scoped>
.ms-content {
    padding: 30px 30px;
}
.text-overflow{
    overflow: hidden;
    text-overflow:ellipsis;
    white-space: nowrap;
}
.btn{
    width:100%
}
a{
    color: #3498db;
}

.explorer{}
.explorer >>> .el-card__header{padding-bottom: 5px;padding-left:10px;}
.explorer >>> .el-card__body{padding:0px}
.bk_icon{
    width:40px;
    height:40px;
    background-color:#f0f0f0;
    text-align: center;
    line-height: 40px;
    font-weight:bold;
    border-radius: 5px;
    float:left
}

.tx_icon{
    border-radius: 20px;
}

.bk_icon_right{
    display:inline-block;
    height:40px;
    margin-left:5px;
}

.txhash{
    display:block;
    width:250px;
    overflow:hidden;
    text-overflow: ellipsis;
}
</style>

