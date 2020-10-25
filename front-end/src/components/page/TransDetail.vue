<template>
<div class="tx-detail">
    <p class="title_bn">{{$t('i18n.txDetail')}}</p>
    <el-card shadow="hover" class="card_content"> 
    <div v-for="item in items" v-bind:key="item.id">
        <el-row :gutter="20">
            <el-col :span="8">
                <el-popover
                    placement="top-start"
                    width="200"
                    trigger="hover"
                    :content="item.content">
                    <i class="el-icon-lx-question" slot="reference"></i>
                </el-popover>
                <span>{{item.field}}</span>
            </el-col>
            <el-col :span="16">
                <!-- normal -->
                <span v-if="item.id != 2 && item.id != 3  && item.id != 5  && item.id != 6 && item.id != 8 && item.id != 12">{{item.value}}</span>

                <!-- status -->
                <div v-if="item.id == 2" class="tx-status" :class="{'tx-fail': item.value == 0}"> {{item.value == 1?"Success":"Fail"}}</div>

                <!-- block num -->
                <div v-if ="item.id == 3">
                    <router-link :to = "'/block/' + item.value"  >{{item.value}}</router-link>
                </div>

                <!-- block num -->
                <div v-if ="item.id == 5 || item.id == 6">
                    <router-link :to = "'/address/' + item.value + '/txs'"  >{{item.value}}</router-link>
                </div>

                <!-- gas used -->
                <div v-if ="item.id == 8">
                     <span>{{item.value}}</span>
                    <span>{{'(' + (tx.gas_limit == 0?0:(item.value/tx.gas_limit*100).toFixed(4)) + '%)'}}</span>
                    <el-progress :show-text="false" :stroke-width="5" :percentage="tx.gas_limit == 0?0:(item.value/tx.gas_limit*100)"></el-progress>
                </div>

                <!-- input data -->
                <div v-if="item.id == 12">
                    <el-input type="textarea" :rows="2" v-model="item.value" disabled></el-input>
                </div>
            </el-col>
        </el-row>
        <hr class ="hr_space">
    </div>
    
    </el-card>
</div>
</template>

<script>
import {dataUrl} from '../../config'
export default {
     data: function() {
        return {
            tx: {
                "tx_hash":"", 
                "block_height":0, 
                "timestamp":0, 
                "from":"", 
                "to":"",
                "gas_limit":1, 
                "gas_price":0,
                "nonce":0, 
                "input":"", 
                "tx_type":2, 
                "value":0, 
                "receipt":{
                    "contract_address":"", 
                    "status":1, 
                    "event":"", 
                    "gas_used":0, 
                }
            }
        }
     },
    props: ['txHash'],
    watch: {
        '$route' (to, from) {
            this.getData()
        }
    },
    methods:{
        getData(){
            this.axios.get(dataUrl + '/tx/' + this.txHash)
            .then(res => {
                this.tx = res.data;
                let d = ["{'From':'三花智控','To':'特斯拉'，'Type':'空调'，'Count':'5','Date':'2020-10-25'}",
                         "{'From':'旭升股份','To':'上汽通用'，'Type':'变速箱箱体'，'Count':'3','Date':'2020-10-25'}",
                         "{'From':'华域汽车','To':'上汽大众'，'Type':'座椅'，'Count':'13','Date':'2020-10-25'}",
                         "{'From':'上海博迩森','To':'上汽大众'，'Type':'内饰'，'Count':'13','Date':'2020-10-25'}",
                         "{'From':'江苏韩泰','To':'上汽大众'，'Type':'轮胎'，'Count':'22','Date':'2020-10-25'}",
                         "{'From':'动睦股份','To':'特斯拉'，'Type':'逆变器'，'Count':'12','Date':'2020-10-25'}",
                         ]
                this.tx.input = d[this.tx.block_height%d.length]
                console.log(this.tx)
            }).catch(res => {
                console.log(res);
            });
        }
    },
    computed:{
        items : function(){
           return  [
                {
                    id:1,
                    content:this.$t('i18n.txhashIndro'),
                    field:"Transaction Hash:",
                    value:this.tx.tx_hash
                },
                {
                    id:2,
                    content:this.$t('i18n.txstatusIndro'),
                    field:"Status:",
                    value:this.tx.receipt.status
                },
                {
                    id:3,
                    content:this.$t('i18n.txblockIndro'),
                    field:"Block:",
                    value:this.tx.block_height
                },
                {
                    id:4,
                    content:this.$t('i18n.txtIndro'),
                    field:"Timestamp:",
                    value:(new Date(this.tx.timestamp * 1000)).toLocaleString()
                },
                {
                    id:5,
                    content:this.$t('i18n.txfromIndro'),
                    field:"From:",
                    value:this.tx.from
                },
                {
                    id:6,
                    content:this.$t('i18n.txtoIndro'),
                    field:"To:",
                    value:this.tx.to
                },
                {
                    id:7,
                    content:this.$t('i18n.txvalueIndro'),
                    field:"Value:",
                    value:this.tx.value
                },
                {
                    id:8,
                    content:this.$t('i18n.txgasusedIndro'),
                    field:"Gas Used:",
                    value:this.tx.receipt.gas_used
                },
                {
                    id:9,
                    content:this.$t('i18n.txgasLimitIndro'),
                    field:"Gas Limit:",
                    value:this.tx.gas_limit
                },
                {
                    id:10,
                    content:this.$t('i18n.txgasPriceIndro'),
                    field:"Gas Price:",
                    value:this.tx.gas_price
                },
                {
                    id:11,
                    content:this.$t('i18n.txnonceIndro'),
                    field:"Nonce:",
                    value:this.tx.nonce
                },
                {
                    id:12,
                    content:this.$t('i18n.txinputIndro'),
                    field:"Input Data:",
                    value:this.tx.input
                }
            ]
        }
    },
    mounted:function() {
        this.getData()
    }
}
</script>
<style scoped>
.title_bn{
    color:#6c757e;
    line-height:1.7;
    margin-top:10px;
    margin-left:10px;
}
.card_content{
    margin-top:20px;
    margin-right:10px;
}
.hr_space{
    margin-top: .75rem;
    margin-bottom: .75rem;
    opacity: .75;
    border-top:1px solid #e7eaf3;
    border-bottom:0px;
    border-left: 0px;
    border-right: 0px;
}
.icon-li-content{
    display: flex;
    height: 100%;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    cursor: pointer;
}

.tx-status{
    color:#00c9a7;
    background-color:rgba(0,201,167,.1);
    font-size: small;
    width:70px;
    height:22px;
    text-align: center;
    border-radius: 3px;
    padding-top:5px;
}

.tx-fail{
    color:#de4437;
    background-color: rgba(222,68,55,.1);
}

a{
    color: #3498db;
}

.tx-detail >>> .el-progress-bar__outer{
    display:block;
    width:400px;
}
</style>