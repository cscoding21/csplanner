import { text } from "@sveltejs/kit"

export const formatDate = (input: any) : string =>  {
    if (!input) {
        return ""
    }

    if (input instanceof Date) {
        return input.toLocaleDateString()
    }

    const o = new Date(Date.parse(input.toString()))
    return o.toLocaleDateString()
}


export const formatDateTime = (input: any) : string =>  {
    if (!input) {
        return ""
    }

    if (input instanceof Date) {
        return getDateDisplay(input) + " at " + input.toLocaleTimeString()
    }

    const o = new Date(Date.parse(input.toString()))
    return getDateDisplay(o) + " at " + o.toLocaleTimeString()
}

const getDateDisplay = (input: Date) : string => {
    const today = new Date()
    const yesterday = new Date(new Date().setDate(new Date().getDate() - 1));

    if(input.toLocaleDateString() === today.toLocaleDateString() ) {
        return "Today"
    }

    if(input.toLocaleDateString() === yesterday.toLocaleDateString()) {
        return "Yesterday"
    }

    return input.toLocaleDateString()
}


// Create our number formatter.
export const formatCurrency = new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
  
    // These options are needed to round to whole numbers if that's what you want.
    //minimumFractionDigits: 0, // (this suffices for whole numbers, but will print 2500.10 as $2,500.1)
    maximumFractionDigits: 0, // (causes 2500.99 to be printed as $2,501)
  });


export const formatPercent = new Intl.NumberFormat('default', {
    style: 'percent',
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  });



export const getInitialsFromName = (name:string) => {
    if (!name) {
        return "XX"
    } 

    if (name.indexOf(" ") === -1) {
        if (name.length === 1) {
            return name.substring(0, 1).toUpperCase() + "X"
        }

        return name.substring(0,2).toUpperCase()
    }

    const names = name.split(" ")

    return names[0].substring(0,1).toUpperCase() + names[names.length - 1].substring(0,1).toUpperCase()
}


export const formatToCommaSepList = (list:string[]):string => {
    if (list.length === 0) {
        return ""
    } else if (list.length === 1) {
        return list[0]
    }

    return list.join(', ')
}

export const pluralize = (input:string, count:number):string => {
    if (count === 1) {
        return input
    }   

    if (input.endsWith("y")) {
        return input.substring(0, input.length - 1) + "ies"
    }

    return input + "s"
}

export const truncateText = (text: string, length: number) : string => {
    if (text.length > length) {
        return text.substring(0, length) + "..."
    }

    return text
}