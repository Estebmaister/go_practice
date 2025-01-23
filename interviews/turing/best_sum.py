from functools import lru_cache
from itertools import combinations

def best_sum_with_restrictions(nums):

    nums = sorted(nums, reverse=True)  # Sort to simplify initial checks
    max_sum = 0
    best_combination = []

    # Try all combinations of numbers to find the best valid sum
    for r in range(1, len(nums) + 1):
        for combination in combinations(nums, r):
            current_sum = 0
            valid = True

            # Validate the combination by the rule
            for num in sorted(combination):
                if num > current_sum:
                    current_sum += num
                else:
                    valid = False
                    break

            # Check if it's the best sum so far
            if valid and current_sum > max_sum:
                max_sum = current_sum
                best_combination = combination

    return max_sum, best_combination

def best_sum_with_restrictions_memo(nums):
    # Remove duplicates and sort descending
    nums = sorted(set(nums), reverse=True)  

    # Memoization cache to avoid recomputation
    @lru_cache(None)
    def dfs(current_sum):
        best_sum = current_sum
        best_combination = []

        # Try all numbers as candidates
        for num in nums:
            if num > current_sum:  # Ensure restriction is followed
                # Recur with updated sum
                new_sum, new_combination = dfs(current_sum + num)
                if new_sum > best_sum:  # Update if a better sum is found
                    best_sum = new_sum
                    best_combination = [num] + new_combination

        return best_sum, best_combination

    # Start from sum 0 and explore all possibilities
    return dfs(0)


# Memoization cache to avoid recomputation
# @lru_cache(None)
def best_sum_with_restrictions_rec(nums, res = 0):
    if len(nums) == 0:
        return res, []
    # Remove duplicates and sort descending
    nums = sorted(set(nums), reverse=True)  

    best_sum = [res for _ in nums]
    best_combination = [[] for _ in nums]

    # Try all numbers as candidates
    counter = -1
    for num in nums:
        counter += 1
        if num > best_sum[counter]:  # Ensure restriction is followed
            # Recursion with updated sum
            new_sum, new_combination = best_sum_with_restrictions_rec([n for n in nums if n>num], best_sum[counter] + num)
            if new_sum > best_sum[counter]:  # Update if a better sum is found
                best_sum[counter] = new_sum
                best_combination[counter] = [num] + new_combination

    idx = best_sum.index(max(best_sum))
    return best_sum[idx], best_combination[idx]

# Example usage with edge cases
print("Best Sum:      %d, Chosen Numbers: %s " %      best_sum_with_restrictions( [5, 1, 4, 8, 2, 20] )) #35
print("Best Sum Memo: %d, Chosen Numbers: %s " % best_sum_with_restrictions_memo( [5, 1, 4, 8, 2, 20] ))
print("Best Sum Rec:  %s, Chosen Numbers: %s " %  best_sum_with_restrictions_rec( [5, 1, 4, 8, 2, 20] ))
print("Best Sum:      %d, Chosen Numbers: %s " %      best_sum_with_restrictions( [5, 1, 3, 7, 2, 10] )) #19
print("Best Sum Memo: %d, Chosen Numbers: %s " % best_sum_with_restrictions_memo( [5, 1, 3, 7, 2, 10] ))
print("Best Sum Rec:  %s, Chosen Numbers: %s " %  best_sum_with_restrictions_rec( [5, 1, 3, 7, 2, 10] ))
print("Best Sum:      %d, Chosen Numbers: %s " %      best_sum_with_restrictions( [1, 1, 1, 1, 1, 1] )) #1
print("Best Sum Memo: %d, Chosen Numbers: %s " % best_sum_with_restrictions_memo( [1, 1, 1, 1, 1, 1] ))
print("Best Sum Rec:  %s, Chosen Numbers: %s " %  best_sum_with_restrictions_rec( [1, 1, 1, 1, 1, 1] ))
print("Best Sum:      %d, Chosen Numbers: %s " %      best_sum_with_restrictions( [50, 20, 30, 10, 5] )) #95
print("Best Sum Memo: %d, Chosen Numbers: %s " % best_sum_with_restrictions_memo( [50, 20, 30, 10, 5] ))
print("Best Sum Rec:  %s, Chosen Numbers: %s " %  best_sum_with_restrictions_rec( [50, 20, 30, 10, 5] ))