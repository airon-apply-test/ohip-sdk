/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { BlockOwnerType } from './BlockOwnerType';
/**
 * Contains a list of block owners.
 * @export
 * @interface BlockOwnersType
 */
export interface BlockOwnersType {
    /**
     *
     * @type {Array<BlockOwnerType>}
     * @memberof BlockOwnersType
     */
    owner?: Array<BlockOwnerType>;
    /**
     * When this flag is true, the logged in user cannot modify the existing block owners for the current block.
     * @type {boolean}
     * @memberof BlockOwnersType
     */
    lockBlockOwners?: boolean;
    /**
     * When this flag is true, the logged in user cannot modify the existing functionSpaceDetails owners for the current block.
     * @type {boolean}
     * @memberof BlockOwnersType
     */
    lockRoomsOwners?: boolean;
    /**
     * When this flag is true, the logged in user cannot modify the existing catering owners for the current block.
     * @type {boolean}
     * @memberof BlockOwnersType
     */
    lockCateringOwners?: boolean;
}
/**
 * Check if a given object implements the BlockOwnersType interface.
 */
export declare function instanceOfBlockOwnersType(value: object): boolean;
export declare function BlockOwnersTypeFromJSON(json: any): BlockOwnersType;
export declare function BlockOwnersTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockOwnersType;
export declare function BlockOwnersTypeToJSON(value?: BlockOwnersType | null): any;
