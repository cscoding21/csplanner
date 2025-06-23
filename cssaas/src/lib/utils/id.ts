/**
 * generate a DB id
 * @returns a new ID sufficient for databases
 */
export const newID = ():string => {
    const uuid = crypto.randomUUID();
    return uuid
}