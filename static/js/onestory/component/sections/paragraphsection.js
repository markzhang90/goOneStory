var templateParagraph = '';
var ParagraphSection = Vue.extend({
    delimiters : ['${', '}'],
  	template: '<div class="ui clearing segment">\
  					<div class="field ui reply form">\
  						<textarea v-on:input="updateValue()" v-model="para"></textarea>\
  					</div>\
  				</div>',
  	props: [
    	'tar_obj',
    ],
    computed: {

    },
    data: function(){
    	return {
    		para: this.tar_obj.para
    	}
    },
    methods: {
    	updateValue: function(){
    		var _self = this
        	this.tar_obj.para = _self.para
    	}
    }
});
