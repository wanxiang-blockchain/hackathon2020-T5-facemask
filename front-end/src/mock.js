let dataUrl = "http://localhost"
const Mock = require('mockjs')
// 修复mock拷贝原生XHR，默认withCredentials是false导致跨域问题
Mock.XHR.prototype.withCredentials = true
Mock.mock(RegExp(dataUrl + '/blocks'), (req, res) => {

    req = req.url.match(/.+blocks\?page_index=(\d+)&page_size=(\d+)/)
    var b = (parseInt(req[1]) - 1) * parseInt(req[2])
    var e = b + parseInt(req[2])
    var data = [{
            height: '22',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '21',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '20',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '19',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '18',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '17',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '16',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '15',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '14',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '13',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '12',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '11',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '10',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '9',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '8',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '7',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '6',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '5',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '4',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            size:45987,
            gas_used:1234442000,
            timestamp:1598519316,
        }, {
            height: '3',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '2',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
        }, {
            height: '1',
            proposer: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_amount:271,
            gas_limit:10000000000,
            gas_used:1234442,
            size:45987,
            timestamp:1598519316,
    }];

    
    if (b >= data.length){
        return []
    }
    if (e >= data.length){
        e = data.length
    }
    return {"page_index":2,"page_size":24,"total":5000,"items" : data.slice(b,e)}
});

Mock.mock(dataUrl + '/nodes', (req, res) => {
    return [
        {
            name: '2146',
            pub_key:"7b6bf3db68bfae0daeeaab69d3f48ae487dd96e9c26609b71d07bd5a2cf0aa070b25002bf80f7f5dc553d046a8b830fc352e6479670f308a3b653dea06be1908",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: true
        },
        {
            name: '2046',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: false
        },
        {
            name: '1946',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: true
        },
        {
            name: '1846',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: false
        },
        {
            name: '1746',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:0,
            is_alive: false
        },
        {
            name: '1646',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: false
        },
        {
            name: '1546',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: false
        },
        {
            name: '1446',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: false
        },
        {
            name: '1346',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: false
        },
        {
            name: '1246',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: false
        },{
            id: '1146',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: false
        },{
            name: '1046',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: false
        },
        {
            name: '0946',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: false
        },
        {
            name: '0846',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: false
        },
        {
            name: '0746',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: false
        },
        {
            name: '0646',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: true
        },
        {
            name: '0546',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: true
        },
        {
            name: '0446',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: true
        },
        {
            name: '0346',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: true
        },
        {
            name: '0246',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: true
        },
        {
            name: '0146',
            pub_key:"0x123123412412412414",
            desc:"",
            internal_ip:"",
            external_ip:"10.65.125.54",
            rpc_port:"6790",
            p2p_port:"16790",
            type:1,
            is_alive: true
        }];
});

Mock.mock(RegExp(dataUrl + '/txs'), (req, res) => {
    req = req.url.match(/.+txs\?page_index=(\d+)&page_size=(\d+)/)
    var b = (parseInt(req[1]) - 1) * parseInt(req[2])
    var e = b + parseInt(req[2])
    var data =  [{
            tx_hash: 'b650d3cc00d81d16c810770d8209f7a79f451468',
            block_height:20,
            from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            gas_limit:1500000000,
            gas_used:2365145,
            timestamp:34354353,
        },{
            tx_hash: 'b650d3cc00d81d16c810770d8209f7a79f451468',
            block_height:20,
            from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            gas_limit:1500000000,
            gas_used:2365145,
            timestamp:34354353,
        }, 
        {
            tx_hash: 'b650d3cc00d81d16c810770d8209f7a79f451468',
            block_height:20,
            from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            gas_limit:1500000000,
            gas_used:2365145,
            timestamp:34354353,
        },
        {
            tx_hash: 'b650d3cc00d81d16c810770d8209f7a79f451468',
            block_height:20,
            from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            gas_limit:1500000000,
            gas_used:2365145,
            timestamp:34354353,
        },
        {
            tx_hash: 'b650d3cc00d81d16c810770d8209f7a79f451468',
            block_height:20,
            from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            gas_limit:1500000000,
            gas_used:2365145,
            timestamp:34354353,
        },
        {
            tx_hash: 'b650d3cc00d81d16c810770d8209f7a79f451468',
            block_height:20,
            from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            gas_limit:1500000000,
            gas_used:2365145,
            timestamp:34354353,
        },
        {
            tx_hash: 'b650d3cc00d81d16c810770d8209f7a79f451468',
            block_height:20,
            from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            gas_limit:1500000000,
            gas_used:2365145,
            timestamp:34354353,
        },
        {
            tx_hash: 'b650d3cc00d81d16c810770d8209f7a79f451468',
            block_height:20,
            from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            gas_limit:1500000000,
            gas_used:2365145,
            timestamp:34354353,
        },{
            tx_hash: 'b650d3cc00d81d16c810770d8209f7a79f451468',
            block_height:20,
            from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            gas_limit:1500000000,
            gas_used:2365145,
            timestamp:34354353,
        }, {
            tx_hash: 'b650d3cc00d81d16c810770d8209f7a79f451468',
            block_height:20,
            from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            gas_limit:1500000000,
            gas_used:2365145,
            timestamp:34354353,
        }, {
            tx_hash: 'b650d3cc00d81d16c810770d8209f7a79f451468',
            block_height:20,
            from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            gas_limit:1500000000,
            gas_used:2365145,
            timestamp:34354353,
        }, {
            tx_hash: 'b650d3cc00d81d16c810770d8209f7a79f451468',
            block_height:20,
            from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            gas_limit:1500000000,
            gas_used:2365145,
            timestamp:34354353,
        }, {
            tx_hash: 'b650d3cc00d81d16c810770d8209f7a79f451468',
            block_height:20,
            from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            gas_limit:1500000000,
            gas_used:2365145,
            timestamp:34354353,
        }, {
            tx_hash: 'b650d3cc00d81d16c810770d8209f7a79f451468',
            block_height:20,
            from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            gas_limit:1500000000,
            gas_used:2365145,
            timestamp:34354353,
    }]
    if (b >= data.length){
        return []
    }
    if (e > data.length){
        e = data.length
    }
    return {page_index:2,page_size:24,total:5000,items : data.slice(b,e)}
});

Mock.mock(RegExp(dataUrl + '/block[^//]'), (req, res) => {
    req = req.url.match(/.+\/block\?block_height=(.+)/)
    var bkNum = req[1]
    console.log(bkNum)
    var data = {
        "hash":"ab123", //区块hash
        "height":parseInt(bkNum), //区块高度
        "timestamp":34354353, //出块时间，时间戳表示
        "tx_amount":300, //包含了多少条交易
        "proposer":"0xabc123",//出块人地址
        "gas_used":3333434, //消耗了多少gas
        "gas_limit":6666666, //gas上限
        "parent_hash":"abc123", //父区块hash
        "extraData":"", //额外信息
        "size":11
    };   
    return data
});

Mock.mock(RegExp(dataUrl + '/block/.+/txs'), (req, res) => {
    console.log(req)
    req = req.url.match(/.+\/block\/(.+)\/txs/)
    var bkNum = req[1]
    var data = {
        "page_index":2, //页号
        "page_size": 24, //本页包含多少条交易数据
        "total":500000, //总数
        "items": [
            {
                tx_hash: '111111111111',
                block_height:bkNum,
                from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
                to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
                gas_limit:1500000000,
                gas_used:2365145,
                timestamp:34354353,
            },
            {
                tx_hash: '222222222222222',
                block_height:bkNum,
                from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
                to:'0xb650d3cc00d81d16c8items10770d8209f7a79f451468',
                gas_limit:1500000000,
                gas_used:2365145,
                timestamp:34354353,
            },
        ],
    }
    console.log(bkNum)
    
return data
});


Mock.mock(RegExp(dataUrl + '/address/from/.+/txs'), (req, res) => {
    console.log(req)
    req = req.url.match(/.+\/address\/from\/(.+)\/txs/)
    var bkNum = req[1]
    var data = {
        "page_index":2, //页号
        "page_size": 24, //本页包含多少条交易数据
        "total":500000, //总数
        "items": [
            {
                tx_hash: '333333333333333333333',
                block_height:bkNum,
                from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
                to:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
                gas_limit:1500000000,
                gas_used:2365145,
                timestamp:34354353,
            },
            {
                tx_hash: '444444444444444444444444',
                block_height:bkNum,
                from: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
                to:'0xb650d3cc00d81d16c8items10770d8209f7a79f451468',
                gas_limit:1500000000,
                gas_used:2365145,
                timestamp:34354353,
            },
        ],
    }
    console.log(bkNum)
    
return data
});

Mock.mock(RegExp(dataUrl + '/tx/.+'), (req, res) => {
    console.log(req)
    req = req.url.match(/.+\/tx\/(.+)/)
    var txhash = req[1]
    var data =   {
        "tx_hash":"ab123", //交易hash
        "block_height":3345, //所在区块高度
        "timestamp":34354353, //出块时间，时间戳表示
        "from":"0x124abc", //发送者地址
        "to":"0xabc123",//接受者地址
        "gas_limit":232423423, //gas上限
        "gas_price":3432423432,//gas价格，wei为单位
        "nonce":34324, //账户交易中随机值，防重放攻击
        "input":"abcabcfefefe", //交易内容
        "tx_type":2, //交易类型
        "value":3343, //转账金额【预留】
        "receipt":{
            "contract_address":"", //合约地址。如果不是部署合约交易，则此字段为空
            "status":1, //交易状态。1：成功，0:失败
            "event":"", //事件。类型是一个json数组，但是因为事件类型不同，暂时无法定义（TODO）
            "gas_used":32434, //消耗多少gas
        }
    }
    console.log(txhash)
    
return data
});



Mock.mock(RegExp(dataUrl + '/contracts'), (req, res) => {
    req = req.url.match(/.+\/contracts\?page_index=(\d+)&page_size=(\d+)/)
    var b = (parseInt(req[1]) - 1) * parseInt(req[2])
    var e = b + parseInt(req[2])
    var data =  [{
            name: 'PTOKEN',
            address:'b650d3cc00d81d16c810770d8209f7a79f451468',
            creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            timestamp: 444555666
        },{
            name: 'PTOKEN',
            address:'b650d3cc00d81d16c810770d8209f7a79f451468',
            creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            timestamp: 444555666
        }, 
        {
            name: 'PTOKEN',
            address:'b650d3cc00d81d16c810770d8209f7a79f451468',
            creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            timestamp: 444555666
        },
        {
            name: 'PTOKEN',
            address:'b650d3cc00d81d16c810770d8209f7a79f451468',
            creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            timestamp: 444555666
        },
        {
            name: 'PTOKEN',
            address:'b650d3cc00d81d16c810770d8209f7a79f451468',
            creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            timestamp: 444555666
        },
        {
            name: 'PTOKEN',
            address:'b650d3cc00d81d16c810770d8209f7a79f451468',
            creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            timestamp: 444555666
        },
        {
            name: 'PTOKEN',
            address:'b650d3cc00d81d16c810770d8209f7a79f451468',
            creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            timestamp: 444555666
        },
        {
            name: 'PTOKEN',
            address:'b650d3cc00d81d16c810770d8209f7a79f451468',
            creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            timestamp: 444555666
        },{
            name: 'PTOKEN',
            address:'b650d3cc00d81d16c810770d8209f7a79f451468',
            creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            timestamp: 444555666
        }, {
            name: 'PTOKEN',
            address:'b650d3cc00d81d16c810770d8209f7a79f451468',
            creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            timestamp: 444555666
        }, {
            name: 'PTOKEN',
            address:'b650d3cc00d81d16c810770d8209f7a79f451468',
            creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            timestamp: 444555666
        }, {
            name: 'PTOKEN',
            address:'b650d3cc00d81d16c810770d8209f7a79f451468',
            creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            timestamp: 444555666
        }, {
            name: 'PTOKEN',
            address:'b650d3cc00d81d16c810770d8209f7a79f451468',
            creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            timestamp: 444555666
        }, {
            name: 'PTOKEN',
            address:'b650d3cc00d81d16c810770d8209f7a79f451468',
            creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
            tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
            timestamp: 444555666
    }]
    if (b >= data.length){
        return []
    }
    if (e > data.length){
        e = data.length
    }
    return {page_index:2,page_size:10,total:5000,items:data.slice(b,e)}
});



Mock.mock(RegExp(dataUrl + '/contract/'), (req, res) => {
    req = req.url.match(/.+\/contract\/(.+)/)
    var addr = req[1]
    var data = {
        name: 'PTOKEN',
        address:'b650d3cc00d81d16c810770d8209f7a79f451468',
        creator: '0xb650d3cc00d81d16c810770d8209f7a79f451468',
        tx_hash:'0xb650d3cc00d81d16c810770d8209f7a79f451468',
        timestamp: 444555666,
        codes:"0xasdasfafqqqqqqqqqq"
    }  
    return data
});


Mock.mock(RegExp(dataUrl + '/users'), (req, res) => {
    var data = [
        {
            account:'yuzongkun@wxblockchain.com',
            role:'admin',
            date:'2020.08.14'
            
        },
        {
            account:'yuzongkun2@wxblockchain.com',
            role:'admin',
            date:'2020.08.14'
        },
        {
            account:'yuzongkun3@wxblockchain.com',
            role:'admin',
            date:'2020.08.14'
        },
        {
            account:'yuzongkun4@wxblockchain.com',
            role:'admin',
            date:'2020.08.14'
        },
        {
            account:'yuzongkun5@wxblockchain.com',
            role:'admin',
            date:'2020.08.14'
        },
        {
            account:'yuzongkun6@wxblockchain.com',
            role:'admin',
            date:'2020.08.14'
        },
        {
            account:'yuzongkun7@wxblockchain.com',
            role:'admin',
            date:'2020.08.14'
        },
        {
            account:'yuzongkun8@wxblockchain.com',
            role:'admin',
            date:'2020.08.14'
        },
        {
            account:'yuzongkun9@wxblockchain.com',
            role:'admin',
            date:'2020.08.14'
        },
        {
            account:'yuzongkun10@wxblockchain.com',
            role:'admin',
            date:'2020.08.14'
        },
        {
            account:'yuzongkun11@wxblockchain.com',
            role:'admin',
            date:'2020.08.14'
        },
    ]
    return data
});

Mock.mock(RegExp(dataUrl + '/cns'), (req, res) => {
    req = req.url.match(/.+\/cns\?page_index=(\d+)&page_size=(\d+)/)
    var b = (parseInt(req[1]) - 1) * parseInt(req[2])
    var e = b + parseInt(req[2])
    var data = [
        {
            name:'contract1',
            version:'1.0.2',
            addr:'0xas28fc8sdf0asasgbvb56123141',
            infos:[
                {
                    version:'1.0.3',
                    addr:'0xojeie9012u1nbijef09fe0-12312'
                },
                {
                    version:'1.0.2',
                    addr:'0xas28fc8sdf0asasgbvb56123141'
                },
                {
                    version:'1.0.1',
                    addr:'0x128912412kjklnmlknasdpkas0-90'
                },
            ]
        },
        {
            name:'contract1',
            version:'1.0.1',
            addr:'0xas28fc8sdf0asasgbvb56123141',
            infos:[
                {
                    version:'1.0.1',
                    addr:'0xas28fc8sdf0asasgbvb56123141',
                }
            ]
        },
        {
            name:'contract1',
            version:'1.0.1',
            addr:'0xas28fc8sdf0asasgbvb56123141',
            infos:[
                {
                    version:'1.0.1',
                    addr:'0xas28fc8sdf0asasgbvb56123141',
                }
            ]
        },
        {
            name:'contract1',
            version:'1.0.1',
            addr:'0xas28fc8sdf0asasgbvb56123141',
            infos:[
                {
                    version:'1.0.1',
                    addr:'0xas28fc8sdf0asasgbvb56123141',
                }
            ]
        },
        {
            name:'contract1',
            version:'1.0.1',
            addr:'0xas28fc8sdf0asasgbvb56123141',
            infos:[
                {
                    version:'1.0.1',
                    addr:'0xas28fc8sdf0asasgbvb56123141',
                }
            ]
        },
        {
            name:'contract1',
            version:'1.0.1',
            addr:'0xas28fc8sdf0asasgbvb56123141',
            infos:[
                {
                    version:'1.0.1',
                    addr:'0xas28fc8sdf0asasgbvb56123141',
                }
            ]
        },
        {
            name:'contract1',
            version:'1.0.1',
            addr:'0xas28fc8sdf0asasgbvb56123141',
            infos:[
                {
                    version:'1.0.1',
                    addr:'0xas28fc8sdf0asasgbvb56123141',
                }
            ]
        },
        {
            name:'contract1',
            version:'1.0.1',
            addr:'0xas28fc8sdf0asasgbvb56123141',
            infos:[
                {
                    version:'1.0.1',
                    addr:'0xas28fc8sdf0asasgbvb56123141',
                }
            ]
        },
        {
            name:'contract1',
            version:'1.0.1',
            addr:'0xas28fc8sdf0asasgbvb56123141',
            infos:[
                {
                    version:'1.0.1',
                    addr:'0xas28fc8sdf0asasgbvb56123141',
                }
            ]
        },
        {
            name:'contract1',
            version:'1.0.1',
            addr:'0xas28fc8sdf0asasgbvb56123141',
            infos:[
                {
                    version:'1.0.1',
                    addr:'0xas28fc8sdf0asasgbvb56123141',
                }
            ]
        },
        {
            name:'contract1',
            version:'1.0.1',
            addr:'0xas28fc8sdf0asasgbvb56123141',
            infos:[
                {
                    version:'1.0.1',
                    addr:'0xas28fc8sdf0asasgbvb56123141',
                }
            ]
        },
        {
            name:'contract1',
            version:'1.0.1',
            addr:'0xas28fc8sdf0asasgbvb56123141',
            infos:[
                {
                    version:'1.0.1',
                    addr:'0xas28fc8sdf0asasgbvb56123141',
                }
            ]
        },
        {
            name:'contract1',
            version:'1.0.1',
            addr:'0xas28fc8sdf0asasgbvb56123141',
            infos:[
                {
                    version:'1.0.1',
                    addr:'0xas28fc8sdf0asasgbvb56123141',
                }
            ]
        },
        {
            name:'contract1',
            version:'1.0.1',
            addr:'0xas28fc8sdf0asasgbvb56123141',
            infos:[
                {
                    version:'1.0.1',
                    addr:'0xas28fc8sdf0asasgbvb56123141',
                }
            ]
        },
    ]
    if (b >= data.length){
        return []
    }
    if (e > data.length){
        e = data.length
    }
    return {page_index:2,page_size:10,total:5000,items:data.slice(b,e)}
});

Mock.mock(RegExp(dataUrl + '/stats$'), (req, res) => {
    var data = {
        "latest_height":333, //最新区块高度
        "total_tx":555, //交易总数
        "total_contract":66, //合约总数
        "total_node":4, //当前节点总数
    };
    return data
});

Mock.mock(RegExp(dataUrl + '/stats/tx/count'), (req, res) => {
    var data = [
        {"date":"2020:9:10","tx_amount":0},
        {"date":"2020:9:11","tx_amount":1},
        {"date":"2020:9:12","tx_amount":2},
        {"date":"2020:9:13","tx_amount":3},
        {"date":"2020:9:14","tx_amount":4},
        {"date":"2020:9:15","tx_amount":5}
    ];
    return data
});