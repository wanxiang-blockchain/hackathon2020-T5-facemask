import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);

export default new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            redirect: '/dashboard'
        },
        {
            path: '/',
            component: () => import(/* webpackChunkName: "home" */ '../components/common/Home.vue'),
            meta: { title: 'home' },
            children: [
                {
                    path: '/explorer',
                    component: () => import(/* webpackChunkName: "explorer" */ '../components/page/Explorer.vue'),
                    meta: { title: 'i18n.explorer' }
                },
                {
                    path: '/blocks',
                    component: () => import(/* webpackChunkName: "blocks" */ '../components/page/Blocks.vue'),
                    meta: { title: 'i18n.blocks'}
                },
                {
                    path: '/blocks/:creator',
                    component: () => import(/* webpackChunkName: "blocks" */ '../components/page/Blocks.vue'),
                    meta: { title: 'i18n.blocks' }
                },
                {
                    path: '/block/:blockNum',
                    component: () => import(/* webpackChunkName: "block" */ '../components/page/BlockDetail.vue'),
                    meta: { title: 'i18n.bkDetail' },
                    props: true
                },
                {
                    path: '/txs',
                    component: () => import(/* webpackChunkName: "txs" */ '../components/page/Transactions.vue'),
                    meta: { title: 'i18n.txs' }
                },
                {
                    path: '/block/:blockNum/txs',
                    component: () => import(/* webpackChunkName: "bktxs" */ '../components/page/Transactions.vue'),
                    meta: { title: 'i18n.txs' },
                    props: true
                },
                {
                    path: '/address/:addr/txs',
                    component: () => import(/* webpackChunkName: "addrtxs" */ '../components/page/Transactions.vue'),
                    meta: { title: 'i18n.txs' },
                    props: true
                },
                {
                    path: '/tx/:txHash',
                    component: () => import(/* webpackChunkName: "tx" */ '../components/page/TransDetail.vue'),
                    meta: { title: 'i18n.txDetail' },
                    props: true
                },
                {
                    path: '/contracts',
                    component: () => import(/* webpackChunkName: "contracts" */ '../components/page/Contracts.vue'),
                    meta: { title: 'i18n.contracts' }
                },
                {
                    path: '/contract/:addr',
                    component: () => import(/* webpackChunkName: "contract" */ '../components/page/ContractDetail.vue'),
                    meta: { title: 'i18n.scDetail' },
                    props: true
                },
                {
                    path: '/system',
                    component: () => import(/* webpackChunkName: "system" */ '../components/page/SystemConfig.vue'),
                    meta: { title: 'i18n.system',permission:true},
                    props: {showEdit:true}
                },
                {
                    path: '/dashboard',
                    component: () => import(/* webpackChunkName: "dashboard" */ '../components/page/Dashboard.vue'),
                    meta: { title: 'i18n.mainPage' }
                },
                {
                    path: '/404',
                    component: () => import(/* webpackChunkName: "404" */ '../components/page/404.vue'),
                    meta: { title: '404' }
                },
                {
                    path: '/403',
                    component: () => import(/* webpackChunkName: "403" */ '../components/page/403.vue'),
                    meta: { title: '403' }
                }
            ]
        },
        {
            path: '/login',
            component: () => import(/* webpackChunkName: "login" */ '../components/page/Login.vue'),
            meta: { title: '登录' }
        },
        {
            path: '*',
            redirect: '/404'
        }
    ]
});
