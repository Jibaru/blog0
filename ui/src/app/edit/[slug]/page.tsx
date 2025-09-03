'use client';

import { useParams } from 'next/navigation';
import PostEditor from '@/components/PostEditor';

export default function EditPostPage() {
  const params = useParams();
  const slug = params.slug as string;

  return <PostEditor postSlug={slug} />;
}