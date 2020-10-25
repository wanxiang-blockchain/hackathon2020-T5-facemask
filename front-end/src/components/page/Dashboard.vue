<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="8">
                <el-card shadow="hover" class="mgb20" style="height:200px;">
                    <div class="user-info">
                        <img src="../../assets/img/img.jpg" class="user-avator" alt />
                        <div class="user-info-cont">
                            <div class="user-info-name">{{name}}</div>
                            <div>{{$t('i18n.' + role)}}</div>
                        </div>
                    </div>
                    <div class="user-info-list">
                        <div>{{$t('i18n.lastLoginTime')}}：</div>
                        <div>{{$t('i18n.lastLoginIP')}}：</div>
                        
                    </div>
                    <div class="user-info-list">
                        <div style="">{{lastLoginTime}}</div>
                        <div>{{lastLoginIP}}</div>
                    </div>
                </el-card>
                <el-card shadow="hover" style="height:200px;">
                    <div>
                        <schart class="schart" canvasId="line" :options="options"></schart>
                    </div>
                </el-card>
            </el-col>
            <el-col :span="16">
                <el-row :gutter="20" class="mgb20">
                    <el-col :span="8">
                        <el-card shadow="hover" :body-style="{padding: '0px'}">
                            <div class="grid-content grid-con-1">
                                <i class="el-icon-document-copy grid-con-icon"></i>
                                <div class="grid-cont-right">
                                    <div class="grid-num">{{stats.latest_height}}</div>
                                    <div>{{$t('i18n.bkh')}}</div>
                                </div>
                            </div>
                        </el-card>
                    </el-col>
                    <el-col :span="8">
                        <el-card shadow="hover" :body-style="{padding: '0px'}">
                            <div class="grid-content grid-con-2">
                                <i class="el-icon-tickets grid-con-icon"></i>
                                <div class="grid-cont-right">
                                    <div class="grid-num">{{stats.total_tx}}</div>
                                    <div>{{$t('i18n.txcnt')}}</div>
                                </div>
                            </div>
                        </el-card>
                    </el-col>
                    <el-col :span="8">
                        <el-card shadow="hover" :body-style="{padding: '0px'}">
                            <div class="grid-content grid-con-3">
                                <i class="el-icon-s-order grid-con-icon"></i>
                                <div class="grid-cont-right">
                                    <div class="grid-num">{{stats.total_contract}}</div>
                                    <div>{{$t('i18n.sccnt')}}</div>
                                </div>
                            </div>
                        </el-card>
                    </el-col>
                </el-row>
                <el-card shadow="hover" class="node_list">
                <Nodes></Nodes>
                </el-card>
            </el-col>
        </el-row>
        <Explorer></Explorer>
    </div>
</template>

<script>
import Schart from 'vue-schart';
import bus from '../common/bus';
import Explorer from './Explorer';
import Nodes from './Nodes';
import {dataUrl} from '../../config'
export default {
    name: 'dashboard',
    data() {
        return {
            name: localStorage.getItem('ms_username'),
            stats:{
                latest_height:0,
                total_tx:0,
                total_contract:0,
                total_node:0
            },
            txStats:{
                labels:[],
                data:[],
            },
            lastLoginTime:'',
            lastLoginIP:''
        };
    },
    methods:{
        getData(){
            this.axios.get(dataUrl + '/stats')
            .then(res => {
                this.stats = res.data;
            }).catch(res => {
                console.log(res);
            });
        },
        getTXStatData(){
            this.axios.get(dataUrl + '/stats/tx/count',{
                params:{
                    num:7
                }
            })
            .then(res => {
                console.log(res)
                for(let i=0;i<res.data.length;i++){
                    let date = res.data[i].date.split(':')
                    if (date.length < 2){
                        date = date[0]
                    }else{
                       date = date[date.length-2] + '.' +  date[date.length-1]
                    }
                    this.txStats.labels.push(date);
                    this.txStats.data.push(res.data[i].tx_amount);
                }
            }).catch(res => {
                console.log(res);
            });
        }
    },
    components: {
        Explorer,
        Nodes,
        Schart
    },
    computed: {
        role() {
            let userrole = localStorage.getItem('ms_userrole');
            return userrole === 'superadmin' ? 'admin' : 'user';
        },
        options() {
            return {
                type: 'line',
                title: {
                    text: this.$t("i18n.txStatistic")
                },
                bgColor: '#fbfbfb',
                labels: this.txStats.labels,
                datasets: [
                    {
                        label: this.$t("i18n.txcnt"),
                        data: this.txStats.data
                    }
                ]
            }
        },
    },
    mounted:function() {
        this.getData();
        this.getTXStatData();
        this.lastLoginTime = localStorage.getItem('ms_last_login_time');
        this.lastLoginIP = localStorage.getItem('ms_last_login_ip');
        setInterval(this.getData(),2000)
    }
};
</script>

<style scoped>
.el-row {
    margin-bottom: 20px;
}

.grid-content {
    display: flex;
    align-items: center;
    height: 70px;
}

.grid-cont-right {
    flex: 1;
    text-align: center;
    font-size: 14px;
    color: #999;
}

.grid-num {
    font-size: 30px;
    font-weight: bold;
}

.grid-con-icon {
    font-size: 50px;
    width: 100px;
    height: 100px;
    text-align: center;
    line-height: 100px;
    color: #fff;
}

.grid-con-1 .grid-con-icon {
    background: rgb(45, 140, 240);
}

.grid-con-1 .grid-num {
    color: rgb(45, 140, 240);
}

.grid-con-2 .grid-con-icon {
    background: rgb(100, 213, 114);
}

.grid-con-2 .grid-num {
    color: rgb(45, 140, 240);
}

.grid-con-3 .grid-con-icon {
    background: rgb(242, 94, 67);
}

.grid-con-3 .grid-num {
    color: rgb(242, 94, 67);
}

.user-info {
    display: flex;
    align-items: center;
    padding-bottom: 20px;
    border-bottom: 2px solid #ccc;
    margin-bottom: 20px;
}

.user-avator {
    width: 70px;
    height: 70px;
    border-radius: 50%;
}

.user-info-cont {
    padding-left: 50px;
    flex: 1;
    font-size: 14px;
    color: #999;
}

.user-info-cont div:first-child {
    font-size: 30px;
    color: #222;
}

.user-info-list {
    font-size: 14px;
    color: #999;
    line-height: 25px;
    display: inline-block;
    margin-left: 30px;
}

.user-info-list span {
    margin-left: 70px;
}

.mgb20 {
    margin-bottom: 20px;
}

.node_list{
    height:330px;
}

.schart {
    width: 100%;
    height: 180px;
}

</style>
