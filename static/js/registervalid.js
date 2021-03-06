$(document).ready(function(){
		$('#registration').validate({
			errorElement:'span',
			errorClass:'help-block',
			focusInvalid:false,
			rules:{
				uname:{
					required:true,
					checkName:true
				},
				pwd:{
					required:true,
				},
				pwd2:{
					required:true,
					equalTo:'#pwd'
				}
			},
			messages:{
				uname:{
					required:"登录名要给一个吧？"
				},
				pwd:{
					required:"小主，密码总要来一个吧？"					
				},
				pwd2:{
					required:"密码要重复一次，免得不记得哦~！",
					equalTo:"两次密码输入不一致，再试一次吧"
				}
			},
			highlight:function(element){
				$(element).closest('.form-group').addClass('has-error');
			},
			success:function(label){
				label.closest('.form-group').removeClass('has-error');
				label.remove();
			},
			errorPlacement:function(error,element){
				element.parent('div').append(error);
			},
			submitHandler:function(form){
				$('#pwd').val(hex_md5($('#pwd').val()));
				$('#pwd2').val(hex_md5($('#pwd2').val()));
				form.submit();
			}
		})
});

jQuery.validator.addMethod("checkName", function(value, element) {
   return this.optional(element) || /^[A-Za-z0-9]+$/.test(value);
}, "用户名只能包括英文字母,数字");
