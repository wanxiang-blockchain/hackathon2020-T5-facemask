<template>
<div class="bk-detail">
    <p class="title_bn">Block #{{blockNum}}</p>
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
                {{item.field}}
            </el-col>
            <el-col :span="16">
                <!-- block id left arrow -->
                <el-button class ="el-icon-d-arrow-left" v-if="item.id==1" size="mini" @click="goLeft"></el-button>

                <!-- normal element -->
                <span v-if ="item.id != 3 && item.id != 4 && item.id != 5 && item.id != 6 && item.id != 8 && item.id != 10">{{item.value}}</span>

                <!-- transaction count -->
                <div v-if ="item.id == 3">
                    <router-link :to = "'/block/' + block.height + '/txs'"  >{{item.value + ' transactions'}}</router-link>
                    <span style="margin-left:5px"> in this block</span>
                </div>

                <!-- mined by -->
                <div v-if ="item.id == 4">
                    <router-link :to = "'/address/' + item.value + '/txs'"  >{{item.value}}</router-link>
                </div>

                <!-- size -->
                <div v-if ="item.id == 5">
                    <span >{{item.value + " bytes"}}</span>
                </div>

                <!-- gas used -->
                <div v-if ="item.id == 6">
                     <span>{{item.value}}</span>
                    <span>{{'('+ (block.gas_limit==0?0:(item.value/block.gas_limit*100)).toFixed(4) + '%)'}}</span>
                    <el-progress :show-text="false" :stroke-width="5" :percentage="block.gas_limit==0?0:(item.value/block.gas_limit*100)"></el-progress>
                </div>

                <!-- extra data -->
                <div v-if="item.id == 8">
                    <el-input type="textarea" :rows="2" v-model="item.value" disabled>
                </el-input>
                </div>

                <!-- parent hash -->
                <div v-if ="item.id == 10">
                    <!-- <router-link :to = "'/block/' + item.value"  >{{item.value}}</router-link> -->
                    <a href="" @click="goParent">{{item.value}}</a>
                </div>

                 <!-- block id right arrow -->
                <el-button class ="el-icon-d-arrow-right" v-if ="item.id==1" size="mini" @click="goRight"></el-button>
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
            block:  {
                "hash":"", 
                "height":0, 
                "timestamp":0, 
                "tx_amount":0, 
                "proposer":"",
                "gas_used":0, 
                "gas_limit":1, 
                "parent_hash":"",
                "extra_data":"", 
                "size":1,
            },
            stats:{
                latest_height:0,
                total_tx:0,
                total_contract:0,
                total_node:0
            }
        } 
     },
    props: ['blockNum'],
    watch: {
        '$route' (to, from) {
            this.getData()
        }
    },
    methods:{
        getData(){
            this.axios.get(dataUrl + '/block',{
                    params: {
                        block_height: this.blockNum
                    }
                }).then(res => {
                this.block = res.data;
                console.log(this.items)
            }).catch(res => {
                console.log(res);
            });
        },
        getStatData(){
            this.axios.get(dataUrl + '/stats')
            .then(res => {
                this.stats = res.data;
            }).catch(res => {
                console.log(res);
            });
        },
        goLeft(){
            if(this.block.height == 1){
                return
            }
            this.$router.push('/block/' + (this.block.height - 1));
        },
        goRight(){
            if(this.block.height >= this.stats.latest_height){
                return
            }
            this.$router.push('/block/' + (this.block.height + 1));
        },
        goParent(e){
            e.preventDefault();
            this.goLeft();
        }
    },
    computed :{
        items:function(){
            return [
                {
                    id:1,
                    content:this.$t('i18n.bkhIndro'),
                    field:"Block Height:",
                    value:this.block.height
                },
                {
                    id:2,
                    content:this.$t('i18n.bktIndro'),
                    field:"Timestamp:",
                    value:(new Date(this.block.timestamp * 1000)).toLocaleString()
                },
                {
                    id:3,
                    content:this.$t('i18n.bktrxIndro'),
                    field:"Transactions:",
                    value:this.block.tx_amount
                },
                {
                    id:4,
                    content:this.$t('i18n.bkminerIndro'),
                    field:"Mined by:",
                    value:this.block.proposer
                },
                {
                    id:5,
                    content:this.$t('i18n.bksizeIndro'),
                    field:"Size:",
                    value:this.block.size
                },
                {
                    id:6,
                    content:this.$t('i18n.bkgasUesdIndro'),
                    field:"Gas Used:",
                    value:this.block.gas_used
                },
                {
                    id:7,
                    content: this.$t('i18n.bkgasLimitIndro'),
                    field:"Gas Limit:",
                    value:this.block.gas_limit
                },
                {
                    id:8,
                    content: this.$t('i18n.bkextraIndro'),
                    field:"Extra Data:",
                    value:this.block.extra_data
                },
                {
                    id:9,
                    content: this.$t('i18n.bkhashIndro'),
                    field:"Hash:",
                    value:this.block.hash
                },
                {
                    id:10,
                    content:this.$t('i18n.bkparentHashIndro'),
                    field:"Parent Hash:",
                    value:this.block.parent_hash
                }
            ]
        }
    },
    mounted:function() {
        this.getData();
        this.getStatData();
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

.el-icon-d-arrow-left{
    color:#3498db;
    margin-right: 10px;
}
.el-icon-d-arrow-left:hover{
    color:crimson
}

.el-icon-d-arrow-right{
    margin-left: 10px;
     color:#3498db;
}

a{
    color: #3498db;
}

.bk-detail >>> .el-progress-bar__outer{
    display:block;
    width:400px;
}
</style>