<template>
<div>
    <p class="title_bn">Contract {{addr}}</p>
    <el-card shadow="hover" class="card_content"> 
    <div v-for="item in items" v-bind:key="item.id">
        <el-row :gutter="20">
            <el-col :span="4">
                <el-popover
                    placement="top-start"
                    width="200"
                    trigger="hover"
                    :content="item.content">
                    <i class="el-icon-lx-question" slot="reference"></i>
                </el-popover>
                {{item.field}}
            </el-col>

            <el-col :span="20" >
                <!-- normal -->
                <span v-if = "item.id != 2 && item.id != 3">{{item.value}}</span>     

                <!-- contract creator -->
                <div v-if = "item.id == 2">
                    <router-link :to = "'/address/' + item.value + '/txs'" >{{item.value}}</router-link>
                    <span style="padding-left:10px;padding-right:10px">at txn</span>
                    <router-link :to = "'/tx/' + item.value2"  >{{item.value2}}</router-link>
                </div>

                <!-- code -->
                <div v-if = "item.id == 3">
                    <el-input type="textarea" :rows="8" v-model="item.value" disabled></el-input>
                </div>
            </el-col>
        </el-row>
        <hr class ="hr_space">
    </div>
    <div>
        <el-button type="primary" @click="froze" v-if="isAdmin">{{$t('i18n.frozeSC')}}</el-button>
    </div>
    </el-card>
</div>
</template>

<script>
import {dataUrl,apiServerUrl} from '../../config'
export default {
     data: function() {
        return {
            contract:{
                "name":"",
                "address":"",
                "creator":"",
                "tx_hash":"",
                "timestamp":"",
                "code":""
            }
        } 
     },
    props: ['addr'],
    watch: {
        '$route' (to, from) {
            this.getData()
        }
    },
    methods:{
        getData(){
            this.axios.get(dataUrl + '/contract/' + this.addr,
            ).then(res => {
                this.contract = res.data;
                console.log(this.items)
            }).catch(res => {
                console.log(res);
            });
        },
        froze(){
            this.$confirm(this.$t('i18n.confirmFroze'), '', {
                confirmButtonText: this.$t('i18n.confirm'),
                cancelButtonText: this.$t('i18n.cancel'),
                type: 'warning'
                }).then(() => {
                    this.doFroze();
                })
        },
        doFroze(){
            this.axios.get(apiServerUrl + '/superadmin/firewall/open',{
                params: {
                    Address:this.contract.address
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
        }
    },
    computed :{
        items:function(){
            return [
                {
                    id:1,
                    content:this.$t('i18n.scNameIndro'),
                    field:"Contract Name:",
                    value:this.contract.name
                },
                {
                    id:2,
                    content:this.$t('i18n.scCreatorIndro'),
                    field:"Contract Creator:",
                    value:this.contract.creator,
                    value2:this.contract.tx_hash
                },
                {
                    id:3,
                    content:this.$t('i18n.scCodeIndro'),
                    field:"Codes:",
                    value:this.contract.code
                }
            ]
        },
        isAdmin:function(){
            let role = localStorage.getItem('ms_userrole');
            return role == "superadmin"
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
    margin-top:20px;
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

a{
    color: #3498db;
}

</style>