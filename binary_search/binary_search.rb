def binary_search(list, target, imin = nil, imax = nil)
  imin ||= 0
  imax ||= list.length - 1
  if(imax < imin)
    return false
  else
    mid = imin + (imax - imin)/2
    if(list[mid] > target)
      return binary_search(list, target, imin, mid -1 )
    elsif(list[mid] < target)
      return binary_search(list, target, mid + 1, imax)
    else
      return mid
    end
  end
end


# ?? (imax + imin) / 2 = imin + (imax - imin) /2    ??
#
# (imax + imin ) /2 = x
# imax + imin = 2x
# imax - imin = 2x - 2imin
# (imax- imin) /2 = x - imin
# imin + (imax - imin) /2 = x
#
#
# (imax + imin) / 2  = imin + (imax - imin) /2
# imax + imin = 2imin + ( imax - imin )
# imax + imin = imax + imin
