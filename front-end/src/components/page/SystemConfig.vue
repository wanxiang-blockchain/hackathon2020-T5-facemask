<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="12">
            <el-card shadow="hover" class="card">
                <Params></Params>
            </el-card>
            </el-col>
            <el-col :span="12">
            <el-card shadow="hover" class="card">
                <Users :height="385"></Users>
            </el-card>
            </el-col>
        </el-row>
        <el-row :gutter="20">
            <el-col :span="12">
                <el-card shadow="hover" class="card">
                    <Nodes :height="385" :showEdit="this.showEdit"></Nodes>
                </el-card>
            </el-col>
            <el-col :span="12">
                <el-card shadow="hover" class="card">
                    <CNSs :height="385"></CNSs>
                </el-card>
            </el-col>
        </el-row>
    </div>
    </template>
</template>

<script>
import Params from './Params';
import Nodes from './Nodes';
import Users from './Users';
import CNSs from './CNSs'
export default {
    props:{
        showEdit :{
            default:false
        }
    },
    components: {
        Params,
        Nodes,
        Users,
        CNSs
    },
}
</script>
<style scoped>
.card{
    height:430px;
    margin-bottom: 15px;
}

.card >>> .el-card__header{padding-bottom: 5px!important;}
</style>