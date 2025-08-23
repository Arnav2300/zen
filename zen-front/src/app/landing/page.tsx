import DotGrid from '@/components/assets/dot-grid';
import { Button } from '@/components/ui/button';
import { SparklesCore } from '@/components/ui/sparkles';
import ZenLogo from '@/components/assets/zen-logo';
import Link from 'next/link';
export default function LandingPage() {
    return (
        // <div style={{ width: '100%', height: '100vh', position: 'relative' }}>
        //     <div className="absolute inset-0 z-0">
        //         <DotGrid
        //             dotSize={5}
        //             gap={15}
        //             baseColor="#061622"
        //             activeColor="#1C9CF0"
        //             proximity={120}
        //             shockRadius={250}
        //             shockStrength={5}
        //             resistance={750}
        //             returnDuration={1.5}
        //         />
        //     </div>
        //     <div className="relative z-10 flex items-center justify-center h-100vh text-white text-center">
        //         Welcome to Zen.
        //     </div>
        //     <div className="relative z-10 flex items-center justify-center h-100vh text-white text-center">
        //         Empower your productivity with Zen — simplicity, speed, and clarity for every workflow.
        //     </div>

        // </div>

        <div className="h-[100vh] w-full bg-black flex flex-col items-center justify-center overflow-hidden rounded-md relative">
            <div className="absolute top-0 left-0 flex justify-between items-center w-full z-30 font-medium p-10">
                <div>
                    <ZenLogo className="text-white w-10 h-10" />
                </div>

                <Link type="button" href="/login">
                    <Button className="font-semibold">Get Started</Button>
                </Link>
            </div>
            <h1 className="md:text-7xl text-3xl lg:text-9xl font-bold text-center text-white relative z-20">
                Zen
            </h1>
            {/* <div className="md:text-3xl text-sm lg:text-2xl font-mono font-medium text-center text-white relative z-20 pb-2">
                Effortless Productivity
            </div> */}
            <div className="w-[40rem] h-40 relative">
                {/* Gradients */}
                <div className="absolute inset-x-20 top-0 bg-gradient-to-r from-transparent via-indigo-500 to-transparent h-[2px] w-3/4 blur-sm" />
                <div className="absolute inset-x-20 top-0 bg-gradient-to-r from-transparent via-indigo-500 to-transparent h-px w-3/4" />
                <div className="absolute inset-x-60 top-0 bg-gradient-to-r from-transparent via-sky-500 to-transparent h-[5px] w-1/4 blur-sm" />
                <div className="absolute inset-x-60 top-0 bg-gradient-to-r from-transparent via-sky-500 to-transparent h-px w-1/4" />

                {/* Core component */}
                <SparklesCore
                    background="transparent"
                    minSize={0.4}
                    maxSize={1}
                    particleDensity={1200}
                    className="w-full h-full"
                    particleColor="#FFFFFF"
                />

                {/* Radial Gradient to prevent sharp edges */}
                <div className="absolute inset-0 w-full h-full bg-black [mask-image:radial-gradient(350px_200px_at_top,transparent_20%,white)]"></div>
            </div>
        </div>
    );



}