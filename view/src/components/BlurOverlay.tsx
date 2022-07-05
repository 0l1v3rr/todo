const BlurOverlay = (props: { isActive: boolean }) => {
    return (
        <div className={`w-screen h-screen absolute top-0 left-0 transition-all duration-300 
            pointer-events-none ${props.isActive ? "backdrop-blur-sm" : ""}`} />
    );
}

export default BlurOverlay