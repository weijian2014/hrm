import type { FormRules } from "element-plus"

export default class Validator {
   /**
    * @description 根据身份证号前17位计算出最后一位
    * @param { string } n 身份证号前17位
    * @returns { string } 最后一位
    */
   private static computeLastNum(n: string): string {
      const computeNums = [7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2]
      const remainderToLast = ["1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"]
      let total = 0
      n.split("").forEach((n, i) => {
         total += Number(n) * computeNums[i]
      })
      return remainderToLast[total % 11]
   }

   /**
    * @description 检验身份证号格式是否正确
    * @param { string } n 身份证号
    * @returns { boolean }
    */
   static isIdentifierValid(n: string): boolean {
      // 检验长度
      if (n.length !== 18) return false
      // 地区是否存在(需要有地区的库，可以查看开源项目)
      // if(!Area.validateAreaCodeExist(n.substr(0, 6))) return false
      // 校验最后一位
      return this.computeLastNum(n.substr(0, 17)) === n.substr(17, 1)
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
   if (h < 120 || h > 230) {
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
   if (w < 40 || w > 130) {
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
   const securityCardString = value?.toString()
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
   if (s < 0) {
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
