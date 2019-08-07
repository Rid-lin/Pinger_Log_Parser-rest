const app = new Vue({
    el: "#app",
    data: {
        editAdress: null,
        adresses: [
            {adress: {
                IP: '',
                Note: '',
                SiteID: '',
                StatusNow:{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },
                StatusOfHour: [
                {
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },{
                    Code: '',
                    Background: '',
                    NumPass:'',
                    NumFail:'',
                },
                ],
            
            },},
        ],
    },
    methods:{
        httpadress(value){
            return `http://${value}`;
        },
        
        deleteAdress(IP){
            i = app.adresses.indexOf(IP)
            fetch("http://127.0.0.1:8080/api/v1/adress"+"/"+IP,{
                method: "DELETE"
            })
            .then(() => {
                delete this.adresses[i]
            })
        },

        updateAdress(adress){
            fetch("http://127.0.0.1:8080/api/v1/adress"+"/"+adress.IP,{
                body: JSON.stringify(adress),
                method: "PUT",
                headers : {"Content-Type": "application/json",},
            }).then(() => {this.editAdress = null;})
        },

        saveConfig(){
            fetch("http://127.0.0.1:8080/api/v1/saveconf")
        },

        createAdress(){
            const adress = {
                IP: '0.0.0.0',
                Note: 'Описание',
                SiteID: 'N\A',
                StatusNow:{
                    Code: '',
                    NumPass: '',
                    NumFail: '',
                },
                StatusOfHour: [
                {
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },{
                    NumPass: '',
                    NumFail: '',
                },
                ],
            };
            this.$set(this.adresses['0.0.0.0']=adress);
            this.editAdress = "0.0.0.0";
        },

        getBackgroundStatusNow(Code) {
            if (Code === '√') return "bggreen"
            if (Code === 'X') return "bgred"
            if (Code === 'O') return "bggrey"
        },

        getBackgroundStatusOfHour(item) {
            if (item.NumFail != 0 && item.NumPass != 0) {
                if (item.NumFail > item.NumPass) {
                    item.Background = "bgpalevioletred"
                } else if (item.NumFail < item.NumPass) {
                    item.Background = "bgyellowgreen"
                } else {
                    item.Background = "bgyellow"
                }
            } else if (item.NumFail == 0 && item.NumPass != 0) {
                item.Background = "bggreen"
            } else if (item.NumFail != 0 && item.NumPass == 0) {
                item.Background = "bgred"
            } else {
                item.Background = "bggrey"
            }
            return item.Background
        },

        getStatusOfHourCode(item) {
            if (item.NumFail != 0 && item.NumPass != 0) {
                if (item.NumFail > item.NumPass) {
                    Code = "X"
                } else if (item.NumFail < item.NumPass) {
                    Code = "√"
                } else {
                    Code = "√"
                }
            } else if (item.NumFail == 0 && item.NumPass != 0) {
                Code = "√"
            } else if (item.NumFail != 0 && item.NumPass == 0) {
                Code = "X"
            } else {
                Code = "O"
            }
            return Code
        },
    },

    mounted(){
        fetch("http://127.0.0.1:8080/api/v1/adresses")
            .then(response => response.json())
            .then((data) => {
                console.log(response)
                this.adresses = data;

            })
    },
})