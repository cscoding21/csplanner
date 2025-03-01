export const particles = () => {
    const canvasElements = document.querySelectorAll('[data-particle-animation]');
    canvasElements.forEach(canvas => {
        const options = {
            //@ts-expect-error
            quantity: canvas.dataset.particleQuantity,
            //@ts-expect-error
            staticity: canvas.dataset.particleStaticity,
            //@ts-expect-error
            ease: canvas.dataset.particleEase,
        };
        new ParticleAnimation(canvas, options);
    });
}

export const highlighter = () => {
    const highlighters = document.querySelectorAll('[data-highlighter]');
    highlighters.forEach((highlighter) => {
        new Highlighter(highlighter);
    });
}