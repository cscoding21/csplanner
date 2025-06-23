import { prisma } from './prisma'
import { newID } from '$lib/utils/id'


export const newEmail = async (email:string) => {
    const id = newID()

    const emailRow = await prisma.email.create({
        data: {
          id: id,
          email: email,
          createdBy: email,
        },
    })
}