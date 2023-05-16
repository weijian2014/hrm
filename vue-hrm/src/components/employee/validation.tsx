import type { FormRules } from "element-plus"

export default class Validator {
   /**
    * @description 检验身份证号格式是否正确
    * @param { string } id 身份证号
    * @returns { boolean }
    */
   static isIdentifierValid(id: string): boolean {
      // 检验长度
      if (id.length !== 18) return false

      var format =
         /^(([1][1-5])|([2][1-3])|([3][1-7])|([4][1-6])|([5][0-4])|([6][1-5])|([7][1])|([8][1-2]))\d{4}(([1][9]\d{2})|([2]\d{3}))(([0][1-9])|([1][0-2]))(([0][1-9])|([1-2][0-9])|([3][0-1]))\d{3}[0-9xX]$/
      //号码规则校验
      if (!format.test(id)) {
         return false
      }
      //区位码校验
      //出生年月日校验   前正则限制起始年份为1900;
      var year = Number(id.substring(6, 10)), //身份证年
         month = Number(id.substring(10, 12)), //身份证月
         date = Number(id.substring(12, 14)), //身份证日
         time = Date.parse(month + "-" + date + "-" + year), //身份证日期时间戳date
         now_time = Date.parse(new Date().toString()), //当前时间戳
         dates = new Date(year, month, 0).getDate() //身份证当月天数
      if (time > now_time || date > dates) {
         return false
      }
      //校验码判断
      var c = new Array(7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2) //系数
      var b = new Array("1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2") //校验码对照表
      var id_array = id.split("")
      var sum = 0
      for (var k = 0; k < 17; k++) {
         sum += parseInt(id_array[k]) * c[k]
      }
      if (id_array[17].toUpperCase() != b[sum % 11].toUpperCase()) {
         return false
      }
      return true
   }

   /**
    * @description 校验手机号格式是否正确
    * @param { string } phone 手机号
    * @returns { boolean }
    */
   static isPhoneValid(phone: string): boolean {
      const phoneRegExp: RegExp =
         /^[1](([3][0-9])|([4][5-9])|([5][0-3,5-9])|([6][5,6])|([7][0-8])|([8][0-9])|([9][1,8,9]))[0-9]{8}$/
      return phoneRegExp.test(phone)
   }
}

const isHeightValid = (rule: unknown, value: string | undefined, callback: (msg?: string) => void) => {
   if (!value) {
      callback("请输入身高, 单位厘米")
      return
   }

   let h = Number(value)
   if (isNaN(h) || h < 120 || h > 230) {
      callback("身高在120~230厘米之间")
      return
   }

   callback()
}

const isWeightValid = (rule: unknown, value: string | undefined, callback: (msg?: string) => void) => {
   if (!value) {
      callback("请输入体重, 单位公斤")
      return
   }

   let w = Number(value)
   if (isNaN(w) || w < 40 || w > 130) {
      callback("体重在40~130公斤之间")
      return
   }

   callback()
}

const isIdentifierValid = (rule: unknown, value: string | undefined, callback: (msg?: string) => void) => {
   if (!value) {
      callback("请输入身份证号")
      return
   }

   const identifierString = value.toString()
   if (identifierString.length != 18) {
      callback("身份证号必须是18位")
      return
   }

   if (!Validator.isIdentifierValid(value?.toString())) {
      callback("身份证号输入有误")
      return
   }

   callback()
}

const isPhoneValid = (rule: unknown, value: string | undefined, callback: (msg?: string) => void) => {
   if (!value) {
      callback("请输入手机号")
      return
   }

   const phoneString = value.toString()
   if (phoneString.length != 11) {
      callback("手机号是11位数字")
      return
   }

   if (!Validator.isPhoneValid(phoneString)) {
      callback("手机号输入有误")
      return
   }

   callback()
}

const isSecurityCardValid = (rule: unknown, value: string | undefined, callback: (msg?: string) => void) => {
   let s = Number(value)
   if (isNaN(s) || s < 0) {
      callback("保安证是6~10位数字")
      return
   }

   const securityCardString = s?.toString()
   if (securityCardString && (securityCardString.length < 6 || securityCardString.length > 10)) {
      callback("保安证是6~10位数字")
      return
   }

   callback()
}

const isSalaryValid = (rule: unknown, value: string | undefined, callback: (msg?: string) => void) => {
   if (!value) {
      callback("请输入工资")
      return
   }

   let s = Number(value)
   if (isNaN(s) || s <= 0) {
      callback("工资是大于等于0的数字, 单位人民币元")
      return
   }

   callback()
}

// 校验规则
export const rules = reactive<FormRules>({
   name: [
      {
         required: true,
         message: "请输入姓名",
         trigger: "blur", // 失去焦点时
      },
   ],
   gender: [
      {
         required: true,
         message: "请选择性别",
         trigger: "blur",
      },
   ],
   birthday: [
      {
         required: true,
         message: "请选择生日",
         trigger: "blur",
      },
   ],
   height: [
      {
         required: true,
         validator: isHeightValid,
         trigger: "blur",
      },
   ],
   weight: [
      {
         required: true,
         validator: isWeightValid,
         trigger: "blur",
      },
   ],
   political_status: [
      {
         required: true,
         message: "请选择政治面貌",
         trigger: "blur",
      },
   ],
   degree: [
      {
         required: true,
         message: "请选择学历",
         trigger: "blur",
      },
   ],
   social_security: [
      {
         required: true,
         message: "请选择社保",
         trigger: "blur",
      },
   ],
   identifier: [
      {
         required: true,
         validator: isIdentifierValid,
         trigger: "blur",
      },
   ],
   phone: [
      {
         required: true,
         validator: isPhoneValid,
         trigger: "blur",
      },
   ],
   first_work_time: [
      {
         required: true,
         message: "请选择首次工作",
         trigger: "blur",
      },
   ],
   post: [
      {
         required: true,
         message: "请输入岗位",
         trigger: "blur",
      },
   ],
   salary: [
      {
         required: true,
         validator: isSalaryValid,
         trigger: "blur",
      },
   ],
   security_card: [
      {
         required: false,
         validator: isSecurityCardValid,
         trigger: "blur",
      },
   ],
})
