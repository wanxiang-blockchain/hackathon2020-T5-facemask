<template>
<div>
    <div slot="header" style="margin-bottom:10px">
        <span>{{$t('i18n.nodeList')}}</span>
        <el-button style="float: right; padding: 3px 10px" type="text" v-if="showEdit" @click="openDialog2">{{$t('i18n.addNodeAccess')}}</el-button>
        <el-button style="float: right; padding: 3px 10px" type="text" v-if="showEdit" @click="openDialog">{{$t('i18n.deployNewNode')}}</el-button>
    </div>

    <el-scrollbar style="height:100%;">
    <el-table :data="nodeList" :max-height="height" style="width:100%" :show-header="false">
    <el-table-column  show-overflow-tooltip>
        <template v-slot="scope" >
            <div class="node_icon"><span>Node</span></div>
            <div class="node_icon_right">
                <div class="node_key">{{scope.row.name + '(' + scope.row.pub_key + ')'}}</div>
                <div style="display:inline-block">{{scope.row.external_ip + ':' + scope.row.p2p_port}}</div>
                <div class = "node_role" v-if="scope.row.type == 1">Validator</div>
            </div>
        </template>
    </el-table-column>
    <el-table-column :width="250" >
            <template v-slot="scope" ><div class = "node_status"  :class="{'node-item-error': !scope.row.is_alive}">{{scope.row.is_alive?"Normal":"Error"}}</div></template>
    </el-table-column>    
    <el-table-column width="120"  v-if="showEdit">
        <template>
            <el-button type="danger" @click="confirm">{{$t('i18n.removeNode')}}</el-button>
        </template>
    </el-table-column>
    <el-table-column type="expand" v-if="showEdit">
        <template slot-scope="props">
            <el-button @click="start(props.row.pub_key)">{{$t('i18n.startNode')}}</el-button>
            <el-button @click="stop(props.row.pub_key)">{{$t('i18n.stopNode')}}</el-button>
            <el-button @click="restart(props.row.pub_key)">{{$t('i18n.restartNode')}}</el-button>
        </template>
    </el-table-column>
    </el-table>
    </el-scrollbar>

    <el-dialog :title="'Node'" :visible.sync="dialogFormVisible">
        <el-form :model="form" label-width="150px"  :label-position="'left'">
            <el-form-item :label="$t('i18n.nodeName') + ':'">
                <el-input v-model="form.nodeName"></el-input>
            </el-form-item>
            <el-form-item :label="$t('i18n.role')">
                <el-select v-model="form.status" placeholder="">
                    <el-option key="Normal" :label="'Normal'" value="Normal"></el-option>
                    <el-option key="Validator" :label="'Validator'" value="Validator"></el-option>
                </el-select>
            </el-form-item>
            <el-form-item :label="$t('i18n.monitorAddr') + ':'">
                <el-input v-model="form.monitorAddr" :placeholder="'address:port'"></el-input>
            </el-form-item>

            <el-form-item :label="'CreatorEnode:'" >
                <el-input v-model="form.creatorEnode" :placeholder="'enode://:pubkey@ip:port;  required if you want to join an exist blockchain'"></el-input>
            </el-form-item>

            <el-form-item :label="'BootNodes:'">
                <el-input type="textarea" :rows="8" v-model="form.bootNodes" :placeholder="'enode://:pubkey@ip:port split by \',\';required if you want to join an exist blockchain'"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="cancel">{{$t('i18n.cancel')}}</el-button>
            <el-button type="primary" @click="submitForm">{{$t('i18n.confirm')}}</el-button>
        </div>
    </el-dialog>


    <el-dialog :title="'Node'" :visible.sync="dialogForm2Visible">
        <el-form :model="form2" label-width="150px"  :label-position="'left'">
            <el-form-item :label="$t('i18n.nodeName') + ':'" >
                <el-input v-model="form2.nodeName"></el-input>
            </el-form-item>
            <el-form-item :label="$t('i18n.role')">
                <el-select v-model="form2.status" placeholder="">
                    <el-option key="Normal" :label="'Normal'" value="Normal"></el-option>
                    <el-option key="Validator" :label="'Validator'" value="Validator"></el-option>
                </el-select>
            </el-form-item>
            <el-form-item :label="'Node Pubkey:'" >
                <el-input v-model="form2.nodePubKey" ></el-input>
            </el-form-item>
            <el-form-item :label="'Node IP:'" >
                <el-input v-model="form2.nodeip" ></el-input>
            </el-form-item>
            <el-form-item :label="'Node P2P Port:'">
                <el-input v-model="form2.p2pPort" ></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="cancel">{{$t('i18n.cancel')}}</el-button>
            <el-button type="primary" @click="submitForm2">{{$t('i18n.confirm')}}</el-button>
        </div>
    </el-dialog>

    
</div>
</template>

<script>
import {dataUrl,apiServerUrl} from '../../config'
export default {
    data() {
        return {
            nodeList: [],
            dialogFormVisible : false,
            form:{
                monitorAddr:'',
                creatorEnode:'',
                bootNodes:'',
                nodeName:'',
                status:'Normal'
            },

            form2:{
                nodeName:'',
                status:'Normal',
                nodeip:'',
                nodePubKey:'',
                p2pPort:''
            }
        }
    },
    methods:{
        getData(){
            this.axios.get(dataUrl + '/nodes').then(res => {
                if (res && res.data) {
                    this.nodeList = res.data;
                }
            }).catch(res => {
                console.log(res)
            })
        },
        confirm(){
            this.$confirm(this.$t('i18n.confirmRemove'), '', {
            confirmButtonText: this.$t('i18n.confirm'),
            cancelButtonText: this.$t('i18n.cancel'),
            type: 'warning'
            }).then(() => {
                this.$message({
                    type: 'success',
                    message: this.$t('i18n.removeSuccess')
                });
            })
        },
        openDialog(){
            this.resetForm()
            this.dialogFormVisible = true
        },
        openDialog2(){
            this.resetForm()
            this.dialogForm2Visible = true
        },
        resetForm(){
            this.form = {
                monitorAddr:"",
                creatorEnode:"",
                bootNodes:"",
                nodeName:'',
                status:'Normal'
            };

            this.form2 = {
                nodeName:'',
                status:'Normal',
                nodeip:'',
                nodePubKey:'',
                p2pPort:''
            }
        },
        submitForm(){
            let monitorInfo = this.form.monitorAddr.split(':')
            let monitorIp = monitorInfo[0]
            let monitorPort = monitorInfo.length > 1?monitorInfo[1]:"50051"
            this.axios.get(apiServerUrl + '/superadmin/createnode',{
                        params: {
                            MonitorIp:monitorIp,
                            MonitorPort:monitorPort,
                            Groupid:0,
                            NodeIp:monitorIp,
                            CreatorEnode:this.form.creatorEnode,
                            Bootnodes:this.form.bootNodes,
                            NodeName:this.form.nodeName,
                            Status:this.form.status == 'Normal'?0:1
                        }
                    }).then(res => {
                if (res && res.data) {
                    console.log(res)
                    this.$message.success("Create Node Success");
                    window.location.reload();
                }
            }).catch(res => {
                this.$message.error("Create Node Failed");
            })
        },
        submitForm2(){
            this.axios.get(apiServerUrl + '/superadmin/node/add',{
                        params: {
                            Name:this.form2.nodeName,
                            ExternalIP:this.form2.nodeip,
                            InternalIP:this.form2.nodeip,
	                        PublicKey:this.form2.nodePubKey,
	                        P2pPort:this.form2.p2pPort,
                            Status:this.form2.status == 'Normal'?0:1
                        }
                    }).then(res => {
                if (res && res.data) {
                    console.log(res)
                    this.$message.success("Create Node Success");
                    window.location.reload();
                }
            }).catch(res => {
                this.$message.error("Create Node Failed");
            })
        },
        cancel(){
            this.resetForm()
            this.dialogFormVisible = false
            this.dialogForm2Visible = false
        },
        start(pub_key){
            this.axios.get(apiServerUrl + '/superadmin/startnode',{
                        params: {
                            PublicKey:pub_key,
                            Groupid:0
                        }
                    }).then(res => {
                if (res && res.data) {
                    console.log(res)
                    this.$message.success("Start Node Success");
                }
            }).catch(res => {
                this.$message.error('Start Node Failed');
            })
        },
        stop(pub_key){
            this.axios.get(apiServerUrl + '/superadmin/stopnode',{
                        params: {
                            PublicKey:pub_key,
                            Groupid:0
                        }
                    }).then(res => {
                if (res && res.data) {
                    console.log(res)
                    this.$message.success("Stop Node Success");
                }
            }).catch(res => {
                this.$message.error('Stop Node Failed');
            })
        },
        restart(pub_key){
            this.axios.get(apiServerUrl + '/superadmin/restartnode',{
                        params: {
                            PublicKey:pub_key,
                            Groupid:0
                        }
                    }).then(res => {
                if (res && res.data) {
                    console.log(res)
                    this.$message.success("Restart Node Success");
                }
            }).catch(res => {
                this.$message.error('Restart Node Failed');
            })
        }
    },
    mounted:function() {
        this.getData()
        setInterval(this.getData,15000)
    },
    props:{
        height:{
            default:270
        },
        showEdit :{
            default:false
        }
    }
}
</script>

<style scoped>

.node_icon{
    width:40px;
    height:40px;
    background-color:#f0f0f0;
    text-align: center;
    line-height: 40px;
    font-weight:bold;
    border-radius: 5px;
    float:left
}

.node_icon_right{
    display:inline-block;
    height:40px;
    margin-left:5px;
}

.node_status{
    width:60px;
    height: 20px;
    background-color: dodgerblue;
    color:white;
    text-align: center;
    border-radius: 3px;
    margin-left: 50px;
}
.node-item-error{
    background-color: crimson;
}
.node_key{
    display: block;
    overflow:hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    max-width: 270px;
}

.node_role{
    width:70px;
    height: 20px;
    background-color: rgba(119,131,143,.1);
    color:#e6a23c;
    text-align: center;
    border-radius: 3px;
    margin-left: 20px;
    font-weight: bold;
    border-top-right-radius: .25rem;
    border-bottom-right-radius : .25rem;
    border-top-left-radius: 0;
    border-bottom-left-radius :0;
    transition: .2s ease-in-out;
    display: inline-block;
}

</style>