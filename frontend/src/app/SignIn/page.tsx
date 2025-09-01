import SigninFormDemo from "@/components/signin-form-demo";
import SignupFormDemo from "@/components/signup-form-demo";
import ThemeToggle from "@/components/ThemeToggle";

export default function SignIn() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center
         bg-gray-100 dark:bg-black transition-colors">
      <div className="absolute top-4 right-4">
        <ThemeToggle />
      </div>
      <SigninFormDemo />
    </div>
  );
}
