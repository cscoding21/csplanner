
/**
 * returns an array of dates normalized to end of week.  The functionality
 * matches the csserver calendar package
 * @param start the begining of the date range
 * @param end the end of the date range
 * @returns an array of week ending dates
 */
export const getCSWeeks = (start:Date, end:Date):Date[] =>  {
    let out:Date[] = []
    let current = endOfWeek(start)

    while(current <= end) {
        out.push(new Date(current))
        current.setDate(current.getDate() + 1)

        current = endOfWeek(current)
    }
	
	return out;
}

/**
 * get the earliest possible date that can be scheduled
 * @returns the earliest possible date that can be scheduled
 */
export const getEarliestScheduleDate = ():Date => {
    const earliest = new Date()
    earliest.setDate(earliest.getDate() + 7)

    return startOfWeek(earliest)
}


function startOfWeek(date:Date):Date {
    var diff = date.getDate() - date.getDay() + (date.getDay() === 0 ? -6 : 1);

    return new Date(date.setDate(diff));
}

function endOfWeek(date:Date):Date {
    var lastday = date.getDate() - (date.getDay() - 1) + 6;
    return new Date(date.setDate(lastday));
}