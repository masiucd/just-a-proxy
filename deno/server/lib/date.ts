export const getDate = (d: string | number = "") => {
  let date = new Date()
  if (d) {
    date = new Date(d)
  }
  const today = date.toLocaleDateString("en-GB", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
  })
  return today
}

export const formatDate = (date: string) => date.replaceAll("/", "-")
export const reverseDate = (date: string) => date.split("-").reverse().join("-")

console.log(reverseDate(formatDate(getDate())))

// 2023-01-21
