var templateImage = '';
var ImageSection = Vue.extend({
    delimiters : ['${', '}'],
  	template: '<div class="ui clearing segment">\
  					<div :id=getImgSectionId class="blurring dimmable image">\
  						<input :id=getImgUploadId type="file" style="display:none" @change="onFileChange">\
  						<div class="ui inverted dimmer">\
  							<div class="content">\
  								<div class="center">\
  									<div class="ui black basic button" @click="uploadImage()">上传图片</div>\
								</div>\
							</div>\
						</div>\
						<img class="ui image" @mouseenter="openDimmer()" @mousedown="closeDimmer()" :src=img_file >\
					</div>\
				</div>',
  	props: [
    	'tar_obj',
    ],
    data: function(){
    	return {
    		img_file: this.tar_obj.upload_img
    	}
    },
    methods: {

    	openDimmer: function(){
            var _self = this;
			$('#'+_self.getImgSectionId).dimmer('show');
    	},
    	closeDimmer: function(){
            var _self = this;
			$('#'+_self.getImgSectionId).dimmer('hide');
    	},
    	uploadImage: function(){
            var _self = this;
    		$('#'+_self.getImgUploadId).click();
    	},
    	onFileChange(e) {
	      	var files = e.target.files || e.dataTransfer.files;
            var _self = this;
	      	if (!files.length)
	        	return;
	      	var _self = this;
			_self.tar_obj.upload_file = files[0]
			
	      	this.createImage(files[0]);
			$('#'+_self.getImgSectionId).dimmer('hide');
	    },
	    createImage(file) {
	      	var image = new Image();
	      	var reader = new FileReader();
	      	var _self = this;

	      	reader.onload = (e) => {
	        	_self.tar_obj.upload_img = e.target.result;
	        	_self.img_file = _self.tar_obj.upload_img;
	      	};
	      	reader.readAsDataURL(file);
	    },
    },
    computed: {
    	getImgSectionId: function(){
            var _self = this;
    		return 'img_section_' + _self.tar_obj.finder;
    	},

    	getImgUploadId: function(){
            var _self = this;
    		return 'img_upload_' + _self.tar_obj.finder;
    	},
    }
});
