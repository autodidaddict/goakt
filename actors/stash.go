/*
 * MIT License
 *
 * Copyright (c) 2022-2024 Tochemey
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package actors

// stash adds the current message to the stash buffer
func (x *pid) stash(ctx ReceiveContext) error {
	if x.stashBuffer == nil {
		return ErrStashBufferNotSet
	}

	x.stashLocker.Lock()
	defer x.stashLocker.Unlock()

	return x.stashBuffer.Push(ctx)
}

// unstash unstashes the oldest message in the stash and prepends to the mailbox
func (x *pid) unstash() error {
	if x.stashBuffer == nil {
		return ErrStashBufferNotSet
	}

	x.stashLocker.Lock()
	defer x.stashLocker.Unlock()

	if !x.stashBuffer.IsEmpty() {
		received, _ := x.stashBuffer.Pop()
		x.doReceive(received)
	}
	return nil
}

// unstashAll unstashes all messages from the stash buffer and prepends in the mailbox
// (it keeps the messages in the same order as received, unstashing older messages before newer).
func (x *pid) unstashAll() error {
	if x.stashBuffer == nil {
		return ErrStashBufferNotSet
	}

	x.stashLocker.Lock()
	defer x.stashLocker.Unlock()

	for !x.stashBuffer.IsEmpty() {
		received, _ := x.stashBuffer.Pop()
		x.doReceive(received)
	}

	return nil
}
