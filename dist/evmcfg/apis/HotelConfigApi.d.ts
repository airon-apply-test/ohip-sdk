/**
 * OPERA Cloud Event Configuration API
 * This API caters for Event Configuration in OPERA Cloud. <br /><There are operations to post, update, fetch and delete codes such as item inventory, function spaces, menu items and many more.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import * as runtime from '../runtime';
import type { ChangeRoomTypePoolRequest, PostRoomTypePoolRequest, RoomTypePoolDetails, Status } from '../models';
export interface ChangeRoomTypePoolOperationRequest {
    roomPoolCode?: string;
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    roomTypePoolToBeChanged?: ChangeRoomTypePoolRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface GetRoomTypePoolRequest {
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    includeInactive?: boolean;
    limit?: number;
    pageNumber?: number;
    physical?: boolean;
    pseudo?: boolean;
    summaryInfo?: boolean;
    roomTypeCodes?: Array<string>;
    roomClassCodes?: Array<string>;
    roomTypePoolCodes?: Array<string>;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface PostRoomTypePoolOperationRequest {
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    roomTypePoolCriteria?: PostRoomTypePoolRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface RemoveRoomTypePoolRequest {
    roomPoolCode?: string;
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    description?: Array<string>;
    defaultText?: string;
    lang?: Array<string>;
    sequence?: Array<number>;
    inactive?: Array<boolean>;
    numberOfRooms?: Array<number>;
    roomType?: Array<string>;
    roomClass?: Array<string>;
    shortDescription?: Array<string>;
    activeDate?: Array<Date>;
    pseudo?: Array<boolean>;
    accessible?: Array<boolean>;
    sendToInterface?: Array<boolean>;
    sellSequence?: Array<number>;
    suite?: Array<boolean>;
    meetingRoom?: Array<boolean>;
    restricted?: Array<boolean>;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
/**
 *
 */
export declare class HotelConfigApi extends runtime.BaseAPI {
    /**
     * Use this API to update Room Type Pool and Associated Room Types. <p><strong>OperationId:</strong>changeRoomTypePool</p>
     * Change Room Type Pool
     */
    changeRoomTypePoolRaw(requestParameters: ChangeRoomTypePoolOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>>;
    /**
     * Use this API to update Room Type Pool and Associated Room Types. <p><strong>OperationId:</strong>changeRoomTypePool</p>
     * Change Room Type Pool
     */
    changeRoomTypePool(requestParameters: ChangeRoomTypePoolOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status>;
    /**
     * This API allows you to Use this API to get Room Type Pool and the associated Room Types. <p><strong>OperationId:</strong>getRoomTypePool</p>
     * Fetch Room Type Pool
     */
    getRoomTypePoolRaw(requestParameters: GetRoomTypePoolRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<RoomTypePoolDetails>>;
    /**
     * This API allows you to Use this API to get Room Type Pool and the associated Room Types. <p><strong>OperationId:</strong>getRoomTypePool</p>
     * Fetch Room Type Pool
     */
    getRoomTypePool(requestParameters: GetRoomTypePoolRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<RoomTypePoolDetails>;
    /**
     * This API allows you to Use this API to create Room Type Pools. <p><strong>OperationId:</strong>postRoomTypePool</p>
     * Create Room Type Pools
     */
    postRoomTypePoolRaw(requestParameters: PostRoomTypePoolOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>>;
    /**
     * This API allows you to Use this API to create Room Type Pools. <p><strong>OperationId:</strong>postRoomTypePool</p>
     * Create Room Type Pools
     */
    postRoomTypePool(requestParameters: PostRoomTypePoolOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status>;
    /**
     * Use this API to delete  Room Type Pool and Mappings. <p><strong>OperationId:</strong>removeRoomTypePool</p>
     * Delete  Room Type Pool and Mappings
     */
    removeRoomTypePoolRaw(requestParameters: RemoveRoomTypePoolRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>>;
    /**
     * Use this API to delete  Room Type Pool and Mappings. <p><strong>OperationId:</strong>removeRoomTypePool</p>
     * Delete  Room Type Pool and Mappings
     */
    removeRoomTypePool(requestParameters: RemoveRoomTypePoolRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status>;
}
