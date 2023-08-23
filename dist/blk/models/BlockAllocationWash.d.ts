/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { BlockAllocationWashType } from './BlockAllocationWashType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request object for block wash operation.
 * @export
 * @interface BlockAllocationWash
 */
export interface BlockAllocationWash {
    /**
     *
     * @type {BlockAllocationWashType}
     * @memberof BlockAllocationWash
     */
    blockAllocationWashType?: BlockAllocationWashType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof BlockAllocationWash
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof BlockAllocationWash
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the BlockAllocationWash interface.
 */
export declare function instanceOfBlockAllocationWash(value: object): boolean;
export declare function BlockAllocationWashFromJSON(json: any): BlockAllocationWash;
export declare function BlockAllocationWashFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockAllocationWash;
export declare function BlockAllocationWashToJSON(value?: BlockAllocationWash | null): any;
