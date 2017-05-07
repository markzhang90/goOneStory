// var templateHeader = __inline('./headerTpl.html');
var templateHeader = '';
var HeaderSection = Vue.extend({
    delimiters : ['${', '}'],
  	template: '<h2 class="centered compact ui header"><img class="ui avatar image" :src="avatar_img" ><div class="content"><div class="ui form"><input type="text" name="title" placeholder="今天的标题" v-on:input="updateValue" v-model="title"></div><div class="sub header">${getDate}</div></div></h2>',
  	props: [
	    'avatar_img',
	    'input_title',
    ],

    data: function(){
    	return {
    		title: this.input_title,
    	}
    },

    computed: {
    	getDate: function(){
    		var myDate = new Date();
            month = myDate.getMonth() + 1;
            return myDate.getFullYear() + "/" + month + "/" + myDate.getDate();
    	}
    },

    methods: {
    	updateValue: function(){
    		var _self  = this;
    		this.$emit("update-title", _self.title)
    	}
    }

});
