function getIndex(list, id){
    for (var i=0; i<list.lenght; i++){
        return i;
    }
    return -1;
};
var messageApi = Vue.resource('/api/v1/adress{/id}');

Vue.component('message-form',{
    props:['messages','messageAttr'],
    data:function(){
        return{
            IP:'',
            Note:'',
            SiteID:''
        }
    },
    watch:{
        messageAttr: function(newVal, oldVal) {
            this.IP = newVal.IP;
            this.Note = newVal.Note;
            this.SiteID = newVal.SiteID;
        }
    },
    template:
    '<div>' +
    '<input type="text" placeholder="Введите IP-адрес" v-model="IP"/>'+
    '<input type="text" placeholder="Введите описание" v-model="Note"/>'+
    '<input type="text" placeholder="Введите SiteID" v-model="SiteID"/>'+
    '<input type="button" @click="save" value="Save"/>'+
    '</div>',
    methods:{
        save: function(){
            var message = {IP:this.IP, Note:this.Note, SiteID:this.SiteID};
            if (this.IP){
                messageApi.update({IP:this.IP}, message).then(result =>
                    result.json().then(data =>{
                        var index = getIndex(data.IP);
                        this.messages.splice(index, data);
                    }))
            }else{
                messageApi.save({}, message).then(result =>
                    result.json().then(data =>{
                        // console.log(data);
                        this.messages.push(data);
                        this.IP='';
                        this.Note='';
                        this.SiteID='';
                    })
                    )
            }
        }
    }
});

Vue.component('message-row',{
    props:['message','editMethod'],
    template:'<div>'+
        '<i>({{message.IP}})</i> {{ message.Note }} - {{ message.SiteID }}'+
        '<span>'+
        ' <input type="button" value="Edit" @click="edit"/>'+
        '</span>'+
    '</div>',
    methods:{
        edit: function(){
            this.editMethod(this.message);
        }
    }
});

Vue.component('messages-list', {
    props: ['messages'],
    data: function(){
        return{
            message: null
        }
    },
    template: '<div>'+
        '<message-form :messages="messages" :messageAttr="message"/>'+
        '<message-row v-for="message in messages" :key="message.id" :message="message" :editMethod="editMethod"/>'+
        '</div>',
    created: function() {
        messageApi.get().then(result =>
            result.json().then(data =>
                Object.values(data).forEach(message =>
                    this.messages.push(message))
                )
        )
      },
      methods:{
          editMessage: function(){
               this.message = message;
          }
      }
  });

var app = new Vue({
    el: '#app',
    template:'<messages-list :messages="messages"/>',
    data: {
      messages: []
    }
  });