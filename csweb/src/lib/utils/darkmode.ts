export const isDarkMode = ():Boolean => {
    const dm = localStorage.getItem("color-theme");

    return (dm === "dark")
}