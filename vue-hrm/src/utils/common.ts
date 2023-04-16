export interface ApiResponse<T> {
   code: number
   message: string
   data: T
}

export type PromiseResponse<T> = Promise<ApiResponse<T>>

export function getAgeByBirthday(stringBirthday: string) {
   let returnAge
   const strBirthdayArr = stringBirthday.split("-")
   const birthYear = Number(strBirthdayArr[0])
   const birthMonth = Number(strBirthdayArr[1])
   const birthDay = Number(strBirthdayArr[2])

   const d = new Date()
   const nowYear = d.getFullYear()
   const nowMonth = d.getMonth() + 1
   const nowDay = d.getDate()

   if (nowYear === birthYear) {
      returnAge = 0 // 同年 则为0岁
   } else {
      const ageDiff = nowYear - birthYear // 年之差
      if (ageDiff > 0) {
         if (nowMonth === birthMonth) {
            const dayDiff = nowDay - birthDay // 日之差
            if (dayDiff < 0) {
               returnAge = ageDiff - 1
            } else {
               returnAge = ageDiff
            }
         } else {
            const monthDiff = nowMonth - birthMonth // 月之差
            if (monthDiff < 0) {
               returnAge = ageDiff - 1
            } else {
               returnAge = ageDiff
            }
         }
      } else {
         returnAge = 0 // 返回-1 表示出生日期输入错误 晚于今天
      }
   }
   return returnAge // 返回周岁年龄
}

export const newDefaultEmployee = (): Employee => {
   const e: Employee = {
      id: 0,
      name: "",
      gender: "",
      birthday: "",
      height: 0,
      weight: 0,
      degree: "",
      identifier: "",
      phone: "",
      political_status: "",
      social_security: "",
      current_address: "",
      first_work_time: "",
      former_employer: "",
      post: "",
      salary: 0,
      security_card: "",
      comments: "",
      updated_at: "",
   }

   return e
}
