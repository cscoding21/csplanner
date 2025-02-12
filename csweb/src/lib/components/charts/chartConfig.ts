import ApexCharts from "apexcharts";

export const getWhiteAndDark900Colors = () => {
  return document.documentElement.classList.contains("dark")
    ? "#1F2937"
    : "#FFFFFF";
}

export const getGray100AndDark700Colors = () => {
  return document.documentElement.classList.contains("dark")
    ? "#374151"
    : "#f3f4f6";
}

export const getPrimaryColor = () => {
  return document.documentElement.classList.contains("dark")
    ? "#1C64F2"
    : "#1A56DB";
}

export const getSecondaryColor = () => {
  return document.documentElement.classList.contains("dark")
    ? "#FF8A4C"
    : "#FF9963";
}

export const getTertiaryColor = () => {
  return document.documentElement.classList.contains("dark")
    ? "#16BDCA"
    : "#12ADBD";
}

export const getQuaternaryColor = () => {
  return document.documentElement.classList.contains("dark")
    ? "#9061F9"
    : "#7E3AF2";
}

export const getQuinaryColor = () => {
  return document.documentElement.classList.contains("dark")
    ? "#F559A5"
    : "#E74694";
}

export const getSuccessColor = () => {
  return document.documentElement.classList.contains("dark")
    ? "#31C48D"
    : "#0E9F6E";
}

export const getDangerColor = () => {
  return document.documentElement.classList.contains("dark")
    ? "#F05252"
    : "#E02424";
}

export const legendColors = {
  class: "text-xs font-normal fill-gray-500 dark:fill-gray-400"
}