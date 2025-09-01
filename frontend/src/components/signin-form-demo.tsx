"use client";

import React from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";

const signupSchema = z
    .object({
        username: z.string().min(3, "Username must be at least 3 characters"),
        password: z.string().min(6, "Password must be at least 6 characters"),
    })


type SignupFormValues = z.infer<typeof signupSchema>;

export default function SigninFormDemo() {
    const {
        register,
        handleSubmit,
        formState: { errors, isSubmitting },
    } = useForm<SignupFormValues>({
        resolver: zodResolver(signupSchema),
    });

    const onSubmit = async (data: SignupFormValues) => {
        const {  ...payload } = data;

        const body = {
            ...payload,
        };

        console.log("Submitting body : ==> ", body);

        try {
            const res = await fetch("http://localhost:3001/SignInUser", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(body),
            });

            if (!res.ok) throw new Error("Failed to create user");
            const responseData = await res.json();
            console.log("✅ User created:", responseData);
            alert("User created successfully!");
        } catch (err) {
            console.error("❌ Error creating user:", err);
            alert("Error creating user");
        }
    };

    return (
        <div className="shadow-input mx-auto w-full max-w-md rounded-none bg-white p-4 md:rounded-2xl md:p-8 dark:bg-zinc-900 dark:border dark:border-zinc-800">
            <h2 className="text-xl font-bold text-neutral-800 dark:text-neutral-100">
                Create your account
            </h2>
            <p className="mt-2 max-w-sm text-sm text-neutral-600 dark:text-neutral-400">
                Fill in your details to sign up
            </p>

            <form className="my-8" onSubmit={handleSubmit(onSubmit)}>
                {/* Username */}
                <LabelInputContainer className="mb-4">
                    <Label
                        htmlFor="username"
                        className="text-neutral-700 dark:text-neutral-200"
                    >
                        Username
                    </Label>
                    <Input
                        id="username"
                        placeholder="jdoe"
                        {...register("username")}
                        className="dark:bg-zinc-800 dark:text-neutral-100 dark:border-zinc-700"
                    />
                    {errors.username && (
                        <p className="text-red-400 text-sm">{errors.username.message}</p>
                    )}
                </LabelInputContainer>




                {/* Password */}
                <LabelInputContainer className="mb-4">
                    <Label
                        htmlFor="password"
                        className="text-neutral-700 dark:text-neutral-200"
                    >
                        Password
                    </Label>
                    <Input
                        id="password"
                        type="password"
                        placeholder="••••••••"
                        {...register("password")}
                        className="dark:bg-zinc-800 dark:text-neutral-100 dark:border-zinc-700"
                    />
                    {errors.password && (
                        <p className="text-red-400 text-sm">{errors.password.message}</p>
                    )}
                </LabelInputContainer>



                <button
                    type="submit"
                    disabled={isSubmitting}
                    className="group/btn relative block h-10 w-full rounded-md bg-gradient-to-br from-black to-neutral-700 font-medium text-white shadow-md dark:from-zinc-800 dark:to-zinc-900"
                >
                    {isSubmitting ? "Signing in..." : "Sign in →"}
                    <BottomGradient />
                </button>
            </form>
        </div>
    );
}

const BottomGradient = () => (
    <>
        <span className="absolute inset-x-0 -bottom-px block h-px w-full bg-gradient-to-r from-transparent via-cyan-500 to-transparent opacity-0 transition duration-500 group-hover/btn:opacity-100" />
        <span className="absolute inset-x-10 -bottom-px mx-auto block h-px w-1/2 bg-gradient-to-r from-transparent via-indigo-500 to-transparent opacity-0 blur-sm transition duration-500 group-hover/btn:opacity-100" />
    </>
);

const LabelInputContainer = ({
    children,
    className,
}: {
    children: React.ReactNode;
    className?: string;
}) => (
    <div className={cn("flex w-full flex-col space-y-2", className)}>
        {children}
    </div>
);
